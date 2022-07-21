package repository

import (
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SongPostgres struct {
	db *sqlx.DB
}

func NewSongPostgres(db *sqlx.DB) *SongPostgres {
	return &SongPostgres{db: db}
}

func (sp *SongPostgres) Create(albumID int, input msh.Song) (msh.Song, error) {
	tx, err := sp.db.Begin()
	if err != nil {
		return msh.Song{}, err
	}

	var newSong msh.Song
	queryAddInSongs := fmt.Sprintf("INSERT INTO %s (title, text, album) VALUES ($1, $2, $3)"+
		" RETURNING id, title, text, album", songsTable)

	row := sp.db.QueryRow(queryAddInSongs, input.Title, input.Text, input.Album)
	if err := row.Scan(&newSong.ID, &newSong.Title, &newSong.Text, &newSong.Album); err != nil {
		_ = tx.Rollback()
		return msh.Song{}, err
	}

	queryAddInAlbumSongs := fmt.Sprintf("INSERT INTO %s (album_id, song_id) VALUES ($1, $2)", albumSongsTable)

	if _, err = sp.db.Exec(queryAddInAlbumSongs, albumID, newSong.ID); err != nil {
		_ = tx.Rollback()
		return msh.Song{}, err
	}

	return newSong, nil
}

func (sp *SongPostgres) GetAll(albumID int) ([]msh.Song, error) {
	tx, err := sp.db.Begin()
	if err != nil {
		return []msh.Song{}, err
	}

	if err = CheckForAvailabilityInAlbums(sp.db, tx, albumID); err != nil {
		return []msh.Song{}, err
	}

	songs := make([]msh.Song, 0)

	query := fmt.Sprintf("SELECT st.id, st.title, st.text, st.album FROM %s st LEFT JOIN %s ast "+
		"on ast.song_id=st.id WHERE ast.album_id=%d", songsTable, albumSongsTable, albumID)

	if err := sp.db.Select(&songs, query); err != nil {
		_ = tx.Rollback()
		return []msh.Song{}, err
	}

	return songs, tx.Commit()
}

func (sp *SongPostgres) GetByID(albumID, songID int) (msh.GetSongOutput, error) {
	tx, err := sp.db.Begin()
	if err != nil {
		return msh.GetSongOutput{}, err
	}

	if err = CheckForAvailabilityInAlbumSongs(sp.db, tx, albumID, songID); err != nil {
		return msh.GetSongOutput{}, err
	}

	var song msh.GetSongOutput
	query := fmt.Sprintf("SELECT st.title, st.text, st.album FROM %s st WHERE st.id=%d", songsTable, songID)

	if err := sp.db.Get(&song, query); err != nil {
		_ = tx.Rollback()
		return msh.GetSongOutput{}, err
	}

	return song, tx.Commit()
}

func (sp *SongPostgres) Delete(albumID, songID int) error {
	tx, err := sp.db.Begin()
	if err != nil {
		return err
	}

	if err = CheckForAvailabilityInAlbumSongs(sp.db, tx, albumID, songID); err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s st WHERE st.id=$1", songsTable)
	if _, err = sp.db.Exec(query, songID); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (sp *SongPostgres) DeleteAll(albumID int) error {
	tx, err := sp.db.Begin()
	if err != nil {
		return err
	}

	if err = DeleteAllSongs(sp.db, tx, albumID); err != nil {
		return err
	}

	return tx.Commit()
}

func (sp *SongPostgres) Update(albumID, songID int, input msh.UpdateSongInput) error {
	tx, err := sp.db.Begin()
	if err != nil {
		return err
	}

	if err = CheckForAvailabilityInAlbumSongs(sp.db, tx, albumID, songID); err != nil {
		return err
	}

	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Text != nil {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argID))
		args = append(args, *input.Text)
		argID++
	}

	if input.Album != nil {
		setValues = append(setValues, fmt.Sprintf("album=$%d", argID))
		args = append(args, *input.Album)
		argID++
	}

	if argID > 1 {
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("UPDATE %s st SET %s WHERE st.id=%d", songsTable, setQuery, songID)

		if _, err = sp.db.Exec(query, args...); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	flag := input.UpdateDate == nil
	queryUpdateDate := fmt.Sprintf("UPDATE %s at SET date=$1 WHERE at.id=$2", albumsTable)

	if !flag {
		if _, err = sp.db.Exec(queryUpdateDate, input.UpdateDate, albumID); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	if input.NewAlbumID != nil {
		if !flag {
			if _, err = sp.db.Exec(queryUpdateDate, input.UpdateDate, input.NewAlbumID); err != nil {
				_ = tx.Rollback()
				return err
			}
		}

		err = CheckForAvailabilityInAlbums(sp.db, tx, *input.NewAlbumID)
		if err != nil {
			_ = tx.Rollback()
			return err
		}

		queryAddNew := fmt.Sprintf("INSERT INTO %s (album_id, song_id) VALUES ($1, $2)", albumSongsTable)
		if _, err = sp.db.Exec(queryAddNew, input.NewAlbumID, songID); err != nil {
			_ = tx.Rollback()
			return err
		}

		queryDelOld := fmt.Sprintf("DELETE FROM %s ast WHERE ast.album_id=$1 AND ast.song_id=$2", albumSongsTable)
		if _, err = sp.db.Exec(queryDelOld, albumID, songID); err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
