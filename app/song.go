package app

type Song struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	Album string `json:"album" binding:"required"`
}
