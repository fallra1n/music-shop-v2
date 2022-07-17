package service

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/repository"
)

type ArtistService struct {
	repo repository.Artist
}

func NewArtistService(repo repository.Artist) *ArtistService {
	return &ArtistService{repo: repo}
}

func (as *ArtistService) CreateArtist(artist msh.Artist) (msh.Artist, error) {
	newArtist, err := as.repo.CreateArtist(artist)
	if err != nil {
		return msh.Artist{}, err
	}
	return newArtist, nil
}
