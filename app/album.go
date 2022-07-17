package app

type Album struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
}
