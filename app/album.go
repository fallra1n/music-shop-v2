package app

type Album struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title" binding:"required"`
	Price  int    `json:"price" db:"price" binding:"required"`
	Artist string `json:"artist" db:"artist" binding:"required"`
	Date   string `json:"date"  db:"date" binding:"required"`
}

type GetAlbum struct {
	Title  string `json:"title" db:"title" binding:"required"`
	Price  int    `json:"price" db:"price" binding:"required"`
	Artist string `json:"artist" db:"artist" binding:"required"`
	Date   string `json:"date"  db:"date" binding:"required"`
}
