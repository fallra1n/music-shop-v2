package service

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/repository"
)

type AlbumService struct {
	repo repository.Album
}

func NewAlbumService(repo repository.Album) *AlbumService {
	return &AlbumService{repo: repo}
}

func (as *AlbumService) Create(artistID int, album msh.Album) (msh.Album, error) {
	newAlbum, err := as.repo.Create(artistID, album)
	if err != nil {
		return msh.Album{}, err
	}
	return newAlbum, nil
}

//func (as *AlbumService) Update(id int, input msh.UpdateArtistInput) error {
//	if err := input.Validate(); err != nil {
//		return err
//	}
//
//	return as.repo.Update(id, input)
//}
//
//func (as *AlbumService) GetAll() ([]msh.Artist, error) {
//	return as.repo.GetAll()
//}
//
func (as *AlbumService) GetByID(artistID, albumID int) (msh.GetAlbum, error) {
	return as.repo.GetByID(artistID, albumID)
}

func (as *AlbumService) DeleteAll(artistID int) error {
	return as.repo.DeleteAll(artistID)
}

//
//func (as *AlbumService) Delete(id int) error {
//	return as.repo.Delete(id)
//}
