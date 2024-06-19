package unit

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Denterry/SocialNetwork/postService/internal/domain/models"
	"github.com/Denterry/SocialNetwork/postService/internal/storage"
	"github.com/google/uuid"
)

func TestGetAuthorIdByPostId(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := storage.NewPostRepositoryPg(db)
	ctx := context.Background()
	postID := int64(1)
	authorUUID := uuid.New()

	rows := sqlmock.NewRows([]string{"author_id"}).AddRow(authorUUID)
	mock.ExpectQuery("SELECT author_id FROM post WHERE id=\\$1").
		WithArgs(postID).
		WillReturnRows(rows)

	authorID, err := repo.GetAuthorIdByPostId(ctx, postID)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if authorID != authorUUID.String() {
		t.Errorf("Expected authorID to be %s, but got %s", authorUUID.String(), authorID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := storage.NewPostRepositoryPg(db)
	ctx := context.Background()
	authorID := uuid.New().String()
	post := models.Post{
		AuthorID: authorID,
		Title:    "Test Title",
		Content:  "Test Content",
	}

	mock.ExpectQuery("INSERT INTO post \\(author_id, title, content\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
		WithArgs(sqlmock.AnyArg(), post.Title, post.Content).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	resultPost, err := repo.CreatePost(ctx, post)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if resultPost.ID != 1 {
		t.Errorf("Expected post ID to be 1, but got %d", resultPost.ID)
	}

	if resultPost.AuthorID != post.AuthorID {
		t.Errorf("Expected authorID to be %s, but got %s", post.AuthorID, resultPost.AuthorID)
	}

	if resultPost.Title != post.Title {
		t.Errorf("Expected title to be %s, but got %s", post.Title, resultPost.Title)
	}

	if resultPost.Content != post.Content {
		t.Errorf("Expected content to be %s, but got %s", post.Content, resultPost.Content)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
