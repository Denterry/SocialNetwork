package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Denterry/SocialNetwork/mainService/domain"
	"github.com/Denterry/SocialNetwork/mainService/util"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user *domain.User)
	ExistsByUsername(username string) bool
	UpdateByUsername(user *domain.User, username string)
	PasswordCheck(username string, password string) uuid.UUID
	Retrieve(username string, password string) *domain.User
	GetUserByID(uid uuid.UUID) *domain.User
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

func (repository *userRepository) PasswordCheck(username string, password string) uuid.UUID {
	// Digest, hash and validation on equelness
	digest := sha256.Sum256([]byte(password))
	user_password := hex.EncodeToString(digest[:])

	// bcrypt?
	// bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	var user *domain.User
	repository.db.Model(&domain.User{}).Where("username = ?", username).Where("password = ?", user_password).First(&user)
	if user == nil {
		return uuid.Nil
	}
	return user.ID
}

func (repository *userRepository) Retrieve(username string, password string) *domain.User {
	user := &domain.User{
		Username: username,
		Password: util.Sha256(password),
	}

	if repository.db.Where(user).First(user).RowsAffected == 0 {
		return nil
	}
	return user
}

func (repository *userRepository) GetUserByID(uid uuid.UUID) *domain.User {
	var user *domain.User
	if err := repository.db.First(&user, uid).Error; err != nil {
		return nil
	}
	user.Password = "***************"
	fmt.Println(user.ID, user.Email, user.Password)
	return user
}
