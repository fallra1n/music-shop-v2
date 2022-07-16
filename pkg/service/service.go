package service

import "github.com/asssswv/music-shop-v2/pkg/repository"

type Artist interface {
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
	return &Service{}
}
