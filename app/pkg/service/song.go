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

func (ss *SongService) GetByID(songID int) (msh.GetSongOutput, error) {
	return ss.repo.GetByID(songID)
}
