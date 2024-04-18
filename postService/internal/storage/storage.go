package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (pr *PostRepositoryPg) GetAuthorByPostId(ctx context.Context, postID int64) (int64, error) {
	query := "SELECT author_id FROM posts  WHERE post_id=$1"
	var authorID int64
	err := pr.db.QueryRowContext(ctx, query, postID).Scan(&authorID)
	if err != nil {
		return -1, err
	}

	return authorID, nil
}

func (pr *PostRepositoryPg) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := "INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3) RETURNING post_id"
	var postID int64
	err := pr.db.QueryRowContext(ctx, query, post.AuthorID, post.Title, post.Content).Scan(&postID)
	if err != nil {
		return nil, err
	}

	return &models.Post{
		ID:       postID,
		AuthorID: post.AuthorID,
		Title:    post.Title,
		Content:  post.Content,
	}, nil
}

func (pr *PostRepositoryPg) UpdatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := "UPDATE posts SET title=$1, content=$2 WHERE post_id=$3"
	_, err := pr.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
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

func (pr *PostRepositoryPg) DeletePost(ctx context.Context, postID int64) error {
	query := "DELETE FROM posts WHERE post_id=$1"
	_, err := pr.db.ExecContext(ctx, query, postID)
	return err
}

func (pr *PostRepositoryPg) GetPostById(ctx context.Context, postID int64) (*models.Post, error) {
	var post models.Post
	query := "SELECT author_id, title, content FROM posts WHERE post_id=$1"
	err := pr.db.QueryRowContext(ctx, query, postID).Scan(&post.AuthorID, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}

	post.ID = postID
	return &post, nil
}

func (pr *PostRepositoryPg) ListPosts(ctx context.Context, pageNumber, pageSize int) ([]*models.Post, error) {
	var posts []*models.Post
	query := fmt.Sprintf("SELECT post_id, author_id, title, content FROM posts LIMIT %d OFFSET %d", pageSize, (pageNumber-1)*pageSize)
	rows, err := pr.db.QueryContext(ctx, query)
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
