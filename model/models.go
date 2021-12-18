package model

// Album represents data about a record Album.
type Album struct {
	ID     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
}