package model

import "github.com/google/uuid"

type UserIdRequest struct {
	UserID uuid.UUID `json:"user_id"`
}
