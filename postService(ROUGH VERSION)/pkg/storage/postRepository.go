package storage

import (
	"context"
	"database/sql"
	"log"

	"github.com/Denterry/SocialNetwork/postService/pkg/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) (int64, error) {
	query := "INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3) RETURNING id"
	var postID int64
	err := r.db.QueryRowContext(ctx, query, post.UserID, post.Title, post.Content).Scan(&postID)
	if err != nil {
		log.Fatal(err)
	}
	return postID, nil
}

// func (r *PostRepository) UpdatePost(ctx context.Context, post *model.Post) error {
// 	query := "UPDATE posts SET title=$1, content=$2 WHERE id=$3"
// 	_, err := r.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
// 	return err
// }

// func (r *PostRepository) DeletePost(ctx context.Context, postID, userID string) error {
// 	query := "DELETE FROM posts WHERE id=$1 AND user_id=$2"
// 	_, err := r.db.ExecContext(ctx, query, postID, userID)
// 	return err
// }

// func (r *PostRepository) GetPostByID(ctx context.Context, postID string) (*model.Post, error) {
// 	var post model.Post
// 	query := "SELECT user_id, title, content FROM posts WHERE id=$1"
// 	err := r.db.QueryRowContext(ctx, query, postID).Scan(&post.UserID, &post.Title, &post.Content)
// 	if err != nil {
// 		return nil, err
// 	}
// 	post.ID = postID
// 	return &post, nil
// }

// func (r *PostRepository) GetPosts(ctx context.Context, pageNumber, pageSize int32) ([]*model.Post, error) {
// 	var posts []*model.Post
// 	query := fmt.Sprintf("SELECT id, user_id, title, content FROM posts LIMIT %d OFFSET %d", pageSize, (pageNumber-1)*pageSize)
// 	rows, err := r.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		var post model.Post
// 		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content); err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, &post)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return posts, nil
// }
