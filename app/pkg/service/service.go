package service

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/repository"
)

type Artist interface {
	CreateArtist(artist msh.Artist) (msh.Artist, error)
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
