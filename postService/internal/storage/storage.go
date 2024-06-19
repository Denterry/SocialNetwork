package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
	"github.com/google/uuid"
)

type postRepositoryPg struct {
	db *sql.DB
}

func NewPostRepositoryPg(db *sql.DB) *postRepositoryPg {
	return &postRepositoryPg{db: db}
}

func (pr *postRepositoryPg) GetAuthorIdByPostId(ctx context.Context, postID int64) (string, error) {
	query := "SELECT author_id FROM post WHERE id=$1"
	var authorID uuid.UUID
	err := pr.db.QueryRowContext(ctx, query, postID).Scan(&authorID)
	if err != nil {
		return "", err
	}

	return authorID.String(), nil
}

func (pr *postRepositoryPg) CreatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	uid, err := uuid.Parse(post.AuthorID) // Parse string to UUID
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO post (author_id, title, content) VALUES ($1, $2, $3) RETURNING id"
	var postID int64
	err = pr.db.QueryRowContext(ctx, query, uid, post.Title, post.Content).Scan(&postID)
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

func (pr *postRepositoryPg) UpdatePost(ctx context.Context, post models.Post) (*models.Post, error) {
	query := "UPDATE post SET title=$1, content=$2 WHERE id=$3"
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

func (pr *postRepositoryPg) DeletePost(ctx context.Context, postID int64) error {
	query := "DELETE FROM post WHERE id=$1"
	_, err := pr.db.ExecContext(ctx, query, postID)
	return err
}

func (pr *postRepositoryPg) GetPostById(ctx context.Context, postID int64) (*models.Post, error) {
	var post models.Post
	var uid uuid.UUID
	query := "SELECT author_id, title, content FROM post WHERE id=$1"
	err := pr.db.QueryRowContext(ctx, query, postID).Scan(&uid, &post.Title, &post.Content)
	if err != nil {
		return nil, err
	}

	post.ID = postID
	post.AuthorID = uid.String()
	return &post, nil
}

func (pr *postRepositoryPg) GetListPosts(ctx context.Context, pageNumber, pageSize int, authorID string) ([]*models.Post, error) {
	var posts []*models.Post
	query := fmt.Sprintf("SELECT id, author_id, title, content FROM post LIMIT %d OFFSET %d", pageSize, (pageNumber-1)*pageSize)
	rows, err := pr.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post models.Post
		var uid uuid.UUID
		if err := rows.Scan(&post.ID, &uid, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		post.AuthorID = uid.String()
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
