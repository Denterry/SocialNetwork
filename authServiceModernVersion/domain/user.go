package domain

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string    `gorm:"type:varchar(100);uniqueIndex"`
	Password string    `gorm:"type:varchar(100);not null"`
	Role     string    `gorm:"type:varchar(100);check: role in ('USER', 'ADMIN')"`
	Name     string    `gorm:"type:varchar(150)"`
	Surname  string    `gorm:"type:varchar(150)"`
	Birthday string    `gorm:"type:varchar(10)"`
	Email    string    `gorm:"uniqueIndex"`
	Phone    string    `gorm:"unique"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	// Digest and store the hex representation.
	digest := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(digest[:])

	return
}
