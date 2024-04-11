package models

type Post struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
