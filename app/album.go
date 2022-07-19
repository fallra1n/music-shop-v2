package app

import "errors"

type Album struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	Price  int    `json:"price" db:"price" binding:"required"`
	Artist string `json:"artist" db:"artist" binding:"required"`
	Date   string `json:"date"  db:"date" binding:"required"`
}

type GetAlbumOutput struct {
	Title  string `json:"title" db:"title" binding:"required"`
	Price  int    `json:"price" db:"price" binding:"required"`
	Artist string `json:"artist" db:"artist" binding:"required"`
	Date   string `json:"date"  db:"date" binding:"required"`
}

type UpdateAlbumInput struct {
	Title      *string `json:"title"`
	Price      *int    `json:"price"`
	Artist     *string `json:"artist"`
	UpdateDate *string `json:"update_date"`
}

func (input *UpdateAlbumInput) Validate() error {
	if input.Title == nil && input.Price == nil && input.Artist == nil && input.UpdateDate == nil {
		return errors.New("update structure has not arguments")
	}

	return nil
}
