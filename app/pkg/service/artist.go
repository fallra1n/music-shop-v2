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

func (as *ArtistService) Create(artist msh.Artist) (msh.Artist, error) {
	newArtist, err := as.repo.Create(artist)
	if err != nil {
		return msh.Artist{}, err
	}
	return newArtist, nil
}

func (as *ArtistService) Update(id int, input msh.UpdateArtistInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return as.repo.Update(id, input)
}

func (as *ArtistService) GetAll() ([]msh.Artist, error) {
	return as.repo.GetAll()
}

func (as *ArtistService) GetByID(id int) (msh.GetArtistWithAlbums, error) {
	return as.repo.GetByID(id)
}

func (as *ArtistService) Delete(id int) error {
	return as.repo.Delete(id)
}
