package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
