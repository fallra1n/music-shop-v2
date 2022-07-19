package app

import "errors"

type Song struct {
	ID    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text" binding:"required"`
	Album string `json:"album" db:"album" binding:"required"`
}

type GetSongOutput struct {
	Title string `json:"title" db:"title" binding:"required"`
	Text  string `json:"text" db:"text" binding:"required"`
	Album string `json:"album"  db:"album" binding:"required"`
}

type UpdateSongInput struct {
	Title      *string `json:"title"`
	Text       *string `json:"artist"`
	Album      *string `json:"update_date"`
	NewAlbumID *int    `json:"new_album_id"`
}

func (input *UpdateSongInput) Validate() error {
	if input.Title == nil && input.Text == nil && input.Album == nil && input.NewAlbumID == nil {
		return errors.New("update structure has not arguments")
	}

	return nil
}
