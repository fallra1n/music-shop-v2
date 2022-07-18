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
}

type Song interface {
}

type Service struct {
	Artist
	Album
	Song
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Artist: NewArtistService(repos.Artist),
	}
}
