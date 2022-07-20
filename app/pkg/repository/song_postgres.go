package repository

import (
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
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
