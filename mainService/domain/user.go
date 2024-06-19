package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"html"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primarykey"`
	Username string    `gorm:"type:varchar(255);uniqueIndex"`
	Password string    `gorm:"type:varchar(100);not null"`
	Role     string    `gorm:"type:varchar(100);check: role in ('USER', 'ADMIN')"`
	Name     string    `gorm:"type:varchar(150)"`
	Surname  string    `gorm:"type:varchar(150)"`
	Birthday string    `gorm:"type:varchar(10)"`
	Email    string    `gorm:"type:varchar(255)"`
	Phone    string    `gorm:"type:varchar(25)"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()

	// Digest and store the hex representation.
	digest := sha256.Sum256([]byte(user.Password))
	user.Password = hex.EncodeToString(digest[:])
	// difference between bcrypt?
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	// user.Password = string(hashedPassword)

	// Remove spaces
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return
}
