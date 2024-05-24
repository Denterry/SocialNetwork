package repository

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User)
	ExistsByUsername(username string) bool
	UpdateByUsername(user *domain.User, username string)
	PasswordCheck(username string, password string) bool
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (repository *userRepository) Save(user *domain.User) {
	repository.db.Save(user)
}

func (repository *userRepository) ExistsByUsername(username string) bool {
	var count int64
	repository.db.Model(&domain.User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

func (repository *userRepository) UpdateByUsername(user *domain.User, username string) {
	repository.db.Model(&domain.User{}).Where("username = ?", username).Updates(user)
}

func (repository *userRepository) PasswordCheck(username string, password string) bool {
	// Digest, hash and validation on equelness
	digest := sha256.Sum256([]byte(password))
	user_password := hex.EncodeToString(digest[:])

	var count int64
	repository.db.Model(&domain.User{}).Where("username = ?", username).Where("password = ?", user_password).Count(&count)
	return count == 1
}
