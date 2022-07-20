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
	songs := make([]msh.Song, 0)

	query := fmt.Sprintf("SELECT st.id, st.title, st.text, st.album FROM %s st LEFT JOIN %s ast "+
		"on ast.song_id=st.id WHERE ast.album_id=%d", songsTable, albumSongsTable, albumID)

	if err := sp.db.Select(&songs, query); err != nil {
		return []msh.Song{}, err
	}

	return songs, nil
}

func (sp *SongPostgres) GetByID(songID int) (msh.GetSongOutput, error) {
	var song msh.GetSongOutput

	query := fmt.Sprintf("SELECT st.title, st.text, st.album FROM %s st WHERE st.id=%d", songsTable, songID)
	if err := sp.db.Get(&song, query); err != nil {
		return msh.GetSongOutput{}, nil
	}

	return song, nil
}
