package models

type Post struct {
	ID       int64  `json:"id"`
	AuthorID int64  `json:"author_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
