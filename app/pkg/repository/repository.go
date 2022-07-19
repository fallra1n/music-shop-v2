package repository

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/jmoiron/sqlx"
)

const (
	artistsTable      = "artists"
	albumsTable       = "albums"
	songsTable        = "songs"
	artistAlbumsTable = "artist_albums"
	albumSongsTable   = "album_songs"
)

type Artist interface {
	Create(artist msh.Artist) (msh.Artist, error)
	Update(id int, input msh.UpdateArtistInput) error
	GetAll() ([]msh.Artist, error)
	GetByID(id int) (msh.GetArtistWithAlbums, error)
	Delete(id int) error
}

type Album interface {
	Create(artistID int, album msh.Album) (msh.Album, error)
	GetByID(artistID, albumID int) (msh.GetAlbumOutput, error)
	DeleteAll(artistID int) error
	Delete(albumID int) error
	Update(albumID int, input msh.UpdateAlbumInput) error
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
		Album:  NewAlbumPostgres(db),
	}
}
