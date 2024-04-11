package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
)

var (
	ErrPostExists   = errors.New("post already exists")
	ErrPostNotFound = errors.New("post not found")
	ErrAppNotFound  = errors.New("app not found")
)

type PostRepositoryPg struct {
	db *sql.DB
}

func NewPostRepositoryPg(db *sql.DB) *PostRepositoryPg {
	return &PostRepositoryPg{db: db}
}

func (pr *PostRepositoryPg) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := "INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3) RETURNING post_id"
	var postID int64
	err := pr.db.QueryRowContext(ctx, query, post.AuthorID, post.Title, post.Content).Scan(&postID)
	if err != nil {
		log.Fatal(err)
	}
	return &models.Post{
		ID:       postID,
		AuthorID: post.AuthorID,
		Title:    post.Title,
		Content:  post.Content,
	}, nil
}

func (r *PostRepositoryPg) UpdatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := "UPDATE posts SET title=$1, content=$2 WHERE id=$3"
	_, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
	if err != nil {
		return nil, err
	}
	return &models.Post{
		ID:       post.ID,
		AuthorID: post.AuthorID,
		Title:    post.Title,
		Content:  post.Content,
	}, nil
}

func (r *PostRepositoryPg) DeletePost(ctx context.Context, postID int64) error {
	query := "DELETE FROM posts WHERE id=$1 AND user_id=$2"
	_, err := r.db.ExecContext(ctx, query, postID)
	return err
}

func (r *PostRepositoryPg) GetPostById(ctx context.Context, postID int64) (*models.Post, error) {
	var post models.Post
	query := "SELECT user_id, title, content FROM posts WHERE id=$1"
	err := r.db.QueryRowContext(ctx, query, postID).Scan(&post.AuthorID, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}
	post.ID = postID
	return &post, nil
}

func (r *PostRepositoryPg) ListPosts(ctx context.Context, pageNumber, pageSize int) ([]*models.Post, error) {
	var posts []*models.Post
	query := fmt.Sprintf("SELECT id, user_id, title, content FROM posts LIMIT %d OFFSET %d", pageSize, (pageNumber-1)*pageSize)
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
