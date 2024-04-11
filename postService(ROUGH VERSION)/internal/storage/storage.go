package storage

import "errors"

var (
	ErrPostExists   = errors.New("post already exists")
	ErrPostNotFound = errors.New("post not found")
	ErrAppNotFound  = errors.New("app not found")
)
