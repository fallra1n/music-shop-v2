package repository

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
)

type Artist interface {
	CreateArtist(artist msh.Artist) (msh.Artist, error)
}

type Album interface {
}

type Song interface {
}

type Repository struct {
	Artist
	Album
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Artist: NewArtistPostgres(db),
	}
}
