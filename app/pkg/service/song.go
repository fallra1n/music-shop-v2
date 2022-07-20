package service

import (
	msh "github.com/asssswv/music-shop-v2/app"
	"github.com/asssswv/music-shop-v2/app/pkg/repository"
)

type SongService struct {
	repo repository.Song
}

func NewSongService(repo repository.Song) *SongService {
	return &SongService{repo: repo}
}

func (ss *SongService) Create(albumID int, input msh.Song) (msh.Song, error) {
	return ss.repo.Create(albumID, input)
}

func (ss *SongService) GetAll(albumID int) ([]msh.Song, error) {
	return ss.repo.GetAll(albumID)
}

func (ss *SongService) GetByID(albumID, songID int) (msh.GetSongOutput, error) {
	return ss.repo.GetByID(albumID, songID)
}

func (ss *SongService) Delete(albumID, songID int) error {
	return ss.repo.Delete(albumID, songID)
}

func (ss *SongService) DeleteAll(albumID int) error {
	return ss.repo.DeleteAll(albumID)
}
