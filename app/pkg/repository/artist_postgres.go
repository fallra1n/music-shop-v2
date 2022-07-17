package repository

import (
	"fmt"
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
)

type ArtistPostgres struct {
	db *sqlx.DB
}

func NewArtistPostgres(db *sqlx.DB) *ArtistPostgres {
	return &ArtistPostgres{db: db}
}

func (ap *ArtistPostgres) CreateArtist(artist msh.Artist) (msh.Artist, error) {
	var id, age int
	var name string
	query := fmt.Sprintf("INSERT INTO %s (name, age) values ($1, $2) RETURNING id, name, age", "artists")
	row := ap.db.QueryRow(query, artist.Name, artist.Age)
	if err := row.Scan(&id, &name, &age); err != nil {
		return msh.Artist{}, err
	}
	return msh.Artist{ID: id, Name: name, Age: age}, nil
}
