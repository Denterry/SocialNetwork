package model

import "github.com/google/uuid"

type CurrentUserRequest struct {
	UserID uuid.UUID `json:"user_id"`
}
