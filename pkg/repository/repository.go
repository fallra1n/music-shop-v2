package repository

import "github.com/jmoiron/sqlx"

type Artist interface {
}

type Album interface {
}

type Song interface {
}

type Repository struct {
	Artist
	Album
	Song
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
