package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Denterry/SocialNetwork/mainService/domain"
	err "github.com/Denterry/SocialNetwork/mainService/error"
	"github.com/Denterry/SocialNetwork/mainService/internal/config"
	"github.com/Denterry/SocialNetwork/mainService/internal/repository"
	"github.com/Denterry/SocialNetwork/mainService/model"
	"github.com/Denterry/SocialNetwork/mainService/util"
	"github.com/google/uuid"
)

type UserService interface {
	Signup(request *model.SignupRequest) error
	ChangeInfo(request *model.ChangeInfoRequest) error
	Signin(request *model.SigninRequest) (string, error)
	Retrieve(request *model.RetrieveRequest) (*model.User, err.ServiceError)
	CurrentUser(request *model.CurrentUserRequest) (*model.User, error)
}

type userService struct {
	repository repository.UserRepository
	cfg        *config.Config
}

func NewUserService(repository repository.UserRepository, cfg *config.Config) UserService {
	return &userService{
		repository: repository,
		cfg:        cfg,
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

func (service *userService) Signin(request *model.SigninRequest) (string, error) {

	exists := service.repository.ExistsByUsername(request.Username)
	if !exists {
		return "", errors.New("user with this username doesn't exist")
	}

	id := service.repository.PasswordCheck(request.Username, request.Password)
	if id == uuid.Nil {
		return "", errors.New("wrong password") // bad practice for hacking
	}

	fmt.Println(fmt.Sprintf("id id id id: %d", id))

	token, err := util.GenerateToken(id, service.cfg)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *userService) Retrieve(request *model.RetrieveRequest) (*model.User, err.ServiceError) {
	user := service.repository.Retrieve(request.Username, request.Password)
	if user == nil {
		return nil, err.NewServiceError(
			"not.found",
			"User not found.",
			http.StatusNotFound,
		)
	}

	return &model.User{Username: user.Username, Role: user.Role}, nil
}

func (service *userService) CurrentUser(request *model.CurrentUserRequest) (*model.User, error) {
	user := service.repository.GetUserByID(request.UserID)
	if user == nil {
		return nil, err.NewServiceError(
			"not.found",
			"User not found.",
			http.StatusNotFound,
		)
	}

	return &model.User{Username: user.Username, Role: user.Role}, nil
}
