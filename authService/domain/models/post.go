package models

type Post struct {
	AuthorID int64  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
