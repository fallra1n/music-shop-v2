package repository

import (
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ArtistPostgres struct {
	db *sqlx.DB
}

func NewArtistPostgres(db *sqlx.DB) *ArtistPostgres {
	return &ArtistPostgres{db: db}
}

func (ap *ArtistPostgres) Create(artist msh.Artist) (msh.Artist, error) {
	var newArtist msh.Artist
	query := fmt.Sprintf("INSERT INTO %s (name, age) values ($1, $2) RETURNING id, name, age", artistsTable)
	row := ap.db.QueryRow(query, artist.Name, artist.Age)
	if err := row.Scan(&newArtist.ID, &newArtist.Name, &newArtist.Age); err != nil {
		return msh.Artist{}, err
	}
	return newArtist, nil
}

func (ap *ArtistPostgres) Update(id int, input msh.UpdateArtistInput) error {
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argID))
		args = append(args, *input.Age)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s a SET %s WHERE a.id = %d", artistsTable, setQuery, id)

	_, err := ap.db.Exec(query, args...)
	return err
}

func (ap *ArtistPostgres) GetAll() ([]msh.Artist, error) {
	artists := make([]msh.Artist, 0)

	query := fmt.Sprintf("SELECT * FROM %s", artistsTable)

	err := ap.db.Select(&artists, query)

	if err != nil {
		return []msh.Artist{}, err
	}

	return artists, nil
}

func (ap *ArtistPostgres) GetByID(id int) (msh.GetArtistWithAlbums, error) {
	var artistWithAlbums msh.GetArtistWithAlbums
	tx, err := ap.db.Begin()

	if err != nil {
		return msh.GetArtistWithAlbums{}, err
	}

	var artist msh.Artist
	getArtistQuery := fmt.Sprintf("SELECT * FROM %s ar WHERE ar.id = %d", artistsTable, id)

	if err = ap.db.Get(&artist, getArtistQuery); err != nil {
		return msh.GetArtistWithAlbums{}, err
	}
	artistWithAlbums.Name = artist.Name
	artistWithAlbums.Age = artist.Age

	getAlbumsIDQuery := fmt.Sprintf("SELECT al.title, al.price, al.artist, al.date FROM %s al LEFT JOIN %s aa "+
		"on aa.album_id=al.id WHERE aa.artist_id=%d", albumsTable, artistAlbumsTable, id)

	if err = ap.db.Select(&artistWithAlbums.Albums, getAlbumsIDQuery); err != nil {
		return msh.GetArtistWithAlbums{}, err
	}

	return artistWithAlbums, tx.Commit()
}

func (ap *ArtistPostgres) Delete(id int) error {
	if err := DeleteAllAlbums(ap.db, id); err != nil {
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s ar WHERE ar.id=$1", artistsTable)
	if _, err := ap.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
