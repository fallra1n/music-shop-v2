package app

import "errors"

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required"`
}

type UpdateArtistInput struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type GetArtistWithAlbums struct {
	Name   string     `json:"name"`
	Age    int        `json:"age"`
	Albums []GetAlbum `json:"albums"`
}

func (input *UpdateArtistInput) Validate() error {
	if input.Name == nil && input.Age == nil {
		return errors.New("update structure has not arguments")
	}

	return nil
}
