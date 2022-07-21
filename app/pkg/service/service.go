package service

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/repository"
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
	Delete(artistID, albumID int) error
	Update(artistID, albumID int, input msh.UpdateAlbumInput) error
}

type Song interface {
	Create(albumID int, input msh.Song) (msh.Song, error)
	GetAll(albumID int) ([]msh.Song, error)
	GetByID(albumID, songID int) (msh.GetSongOutput, error)
	Delete(albumID, songID int) error
	DeleteAll(albumID int) error
	Update(albumID, songID int, input msh.UpdateSongInput) error
}

type Service struct {
	Artist
	Album
	Song
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Artist: NewArtistService(repos.Artist),
		Album:  NewAlbumService(repos.Album),
		Song:   NewSongService(repos.Song),
	}
}
