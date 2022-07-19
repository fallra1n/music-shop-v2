package repository

import (
	"errors"
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
	"strings"
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

func (ap *AlbumPostgres) GetByID(artistID, albumID int) (msh.GetAlbumOutput, error) {
	tx, err := ap.db.Begin()
	if err != nil {
		return msh.GetAlbumOutput{}, err
	}
	queryCheck := fmt.Sprintf("SELECT * FROM %s aa WHERE aa.artist_id=$1 AND aa.album_id=$2", artistAlbumsTable)
	a, er := ap.db.Exec(queryCheck, artistID, albumID)

	if er != nil {
		_ = tx.Rollback()
		return msh.GetAlbumOutput{}, err
	}

	rAff, rAffErr := a.RowsAffected()
	if rAffErr != nil {
		return msh.GetAlbumOutput{}, rAffErr
	}

	if rAff == 0 {
		return msh.GetAlbumOutput{}, errors.New("this album or artist was deleted")
	}

	var album msh.GetAlbumOutput
	queryGet := fmt.Sprintf("SELECT al.title, al.price, al.artist, al.date FROM %s al WHERE al.id=$1", albumsTable)

	if err = ap.db.Get(&album, queryGet, albumID); err != nil {
		_ = tx.Rollback()
		return msh.GetAlbumOutput{}, err
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

func (ap *AlbumPostgres) Delete(albumID int) error {
	query := fmt.Sprintf("DELETE FROM %s al WHERE al.id=$1", albumsTable)

	if _, err := ap.db.Exec(query, albumID); err != nil {
		return err
	}

	return nil
}

func (ap *AlbumPostgres) Update(albumID int, input msh.UpdateAlbumInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	if input.Artist != nil {
		setValues = append(setValues, fmt.Sprintf("artist=$%d", argID))
		args = append(args, *input.Artist)
		argID++
	}

	if input.UpdateDate != nil {
		setValues = append(setValues, fmt.Sprintf("update_date=$%d", argID))
		args = append(args, *input.UpdateDate)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s a SET %s WHERE a.id = %d", albumsTable, setQuery, albumID)

	_, err := ap.db.Exec(query, args...)
	return err
}
