package service

import (
	"errors"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/domain"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/internal/repository"
	"github.com/Denterry/SocialNetwork/authServiceModernVersion/model"
)

type UserService interface {
	Signup(request *model.SignupRequest) error
	ChangeInfo(request *model.ChangeInfoRequest) error
	Signin(request *model.SigninRequest) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (service *userService) Signup(request *model.SignupRequest) error {
	if request.Password != request.PasswordConfirmation {
		return errors.New("password and confirm password must match")
	}

	exists := service.repository.ExistsByUsername(request.Username)
	if exists {
		return errors.New("username already exists")
	}

	service.repository.Save(&domain.User{
		Username: request.Username,
		Password: request.Password,
		Role:     "USER",
	})

	return nil
}

func (service *userService) ChangeInfo(request *model.ChangeInfoRequest) error {
	// here should preferably be validation on access

	exists := service.repository.ExistsByUsername(request.Username)
	if !exists {
		return errors.New("user with this username doesn't exist")
	}

	service.repository.UpdateByUsername(&domain.User{
		Name:     request.Name,
		Surname:  request.Surname,
		Birthday: request.Birthday,
		Email:    request.Email,
		Phone:    request.Phone,
	}, request.Username)

	return nil
}

func (service *userService) Signin(request *model.SigninRequest) error {

	exists := service.repository.ExistsByUsername(request.Username)
	if !exists {
		return errors.New("user with this username doesn't exist")
	}

	access := service.repository.PasswordCheck(request.Username, request.Password)
	if !access {
		return errors.New("wrong password")
	}

	return nil
}
