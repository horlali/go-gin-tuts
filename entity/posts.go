package entity

type Post struct {
	Body  string `json:"body" binding:"required"`
	Title string `json:"title" binding:"required"`
}
