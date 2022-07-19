package repository

import (
	"errors"
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
)

type AlbumPostgres struct {
	db *sqlx.DB
}

func NewAlbumPostgres(db *sqlx.DB) *AlbumPostgres {
	return &AlbumPostgres{db: db}
}

func (ap *AlbumPostgres) Create(artistID int, album msh.Album) (msh.Album, error) {
	tx, err := ap.db.Begin()

	if err != nil {
		return msh.Album{}, err
	}

	var newAlbum msh.Album
	queryAddInAlbums := fmt.Sprintf("INSERT INTO %s (title, price, artist, date) VALUES ($1, $2, $3, $4)"+
		" RETURNING id, title, price, artist, date", albumsTable)

	row := ap.db.QueryRow(queryAddInAlbums, album.Title, album.Price, album.Artist, album.Date)
	if err := row.Scan(&newAlbum.ID, &newAlbum.Title, &newAlbum.Price, &newAlbum.Artist, &newAlbum.Date); err != nil {
		_ = tx.Rollback()
		return msh.Album{}, err
	}

	queryAddInArtistAlbums := fmt.Sprintf("INSERT INTO %s (artist_id, album_id) VALUES ($1, $2)", artistAlbumsTable)

	_, err = ap.db.Exec(queryAddInArtistAlbums, artistID, newAlbum.ID)
	if err != nil {
		_ = tx.Rollback()
		return msh.Album{}, err
	}

	return newAlbum, tx.Commit()
}

func (ap *AlbumPostgres) GetByID(artistID, albumID int) (msh.GetAlbum, error) {
	tx, err := ap.db.Begin()
	if err != nil {
		return msh.GetAlbum{}, err
	}
	queryCheck := fmt.Sprintf("SELECT * FROM %s aa WHERE aa.artist_id=$1 AND aa.album_id=$2", artistAlbumsTable)
	a, er := ap.db.Exec(queryCheck, artistID, albumID)

	if er != nil {
		_ = tx.Rollback()
		return msh.GetAlbum{}, err
	}

	rAff, rAffErr := a.RowsAffected()
	if rAffErr != nil {
		return msh.GetAlbum{}, rAffErr
	}

	if rAff == 0 {
		return msh.GetAlbum{}, errors.New("this album or artist was deleted")
	}

	var album msh.GetAlbum
	queryGet := fmt.Sprintf("SELECT al.title, al.price, al.artist, al.date FROM %s al WHERE al.id=$1", albumsTable)

	if err = ap.db.Get(&album, queryGet, albumID); err != nil {
		_ = tx.Rollback()
		return msh.GetAlbum{}, err
	}
	artistID++
	return album, nil
}

func (ap *AlbumPostgres) DeleteAll(artistID int) error {
	if err := DeleteAll(ap.db, artistID); err != nil {
		return err
	}

	return nil
}
