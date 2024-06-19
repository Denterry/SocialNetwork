package unit

// import (
// 	"testing"

// 	"github.com/Denterry/SocialNetwork/mainService/domain"
// 	"github.com/Denterry/SocialNetwork/mainService/internal/service"
// 	"github.com/Denterry/SocialNetwork/mainService/model"
// 	"github.com/Denterry/SocialNetwork/mainService/util"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // MockUserRepository is a mock of the UserRepository interface
// type MockUserRepository struct {
// 	mock.Mock
// }

// func (m *MockUserRepository) Save(user *domain.User) {
// 	m.Called(user)
// }

// func (m *MockUserRepository) ExistsByUsername(username string) bool {
// 	args := m.Called(username)
// 	return args.Bool(0)
// }

// func (m *MockUserRepository) UpdateByUsername(user *domain.User, username string) {
// 	m.Called(user, username)
// }

// func (m *MockUserRepository) PasswordCheck(username string, password string) uuid.UUID {
// 	args := m.Called(username, password)
// 	return args.Get(0).(uuid.UUID)
// }

// func (m *MockUserRepository) Retrieve(username string, password string) *domain.User {
// 	args := m.Called(username, password)
// 	if args.Get(0) == nil {
// 		return nil
// 	}
// 	return args.Get(0).(*domain.User)
// }

// func (m *MockUserRepository) GetUserByID(uid uuid.UUID) *domain.User {
// 	args := m.Called(uid)
// 	if args.Get(0) == nil {
// 		return nil
// 	}
// 	return args.Get(0).(*domain.User)
// }

// func TestUserService_Signup(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userService := service.NewUserService(mockRepo, nil)

// 	t.Run("successful signup", func(t *testing.T) {
// 		signupRequest := &model.SignupRequest{
// 			Username:             "test@example.com",
// 			Password:             "password123",
// 			PasswordConfirmation: "password123",
// 		}

// 		mockRepo.On("ExistsByUsername", signupRequest.Username).Return(false)
// 		mockRepo.On("Save", mock.AnythingOfType("*domain.User")).Return()

// 		err := userService.Signup(signupRequest)

// 		assert.NoError(t, err)
// 		mockRepo.AssertExpectations(t)
// 	})

// 	t.Run("password confirmation does not match", func(t *testing.T) {
// 		signupRequest := &model.SignupRequest{
// 			Username:             "test@example.com",
// 			Password:             "password123",
// 			PasswordConfirmation: "wrongpassword",
// 		}

// 		err := userService.Signup(signupRequest)

// 		assert.Error(t, err)
// 		assert.Equal(t, "password and confirm password must match", err.Error())
// 	})
// }

// func TestUserService_Signin(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userService := service.NewUserService(mockRepo, nil)

// 	t.Run("successful signin", func(t *testing.T) {
// 		signinRequest := &model.SigninRequest{
// 			Username: "test@example.com",
// 			Password: "password123",
// 		}

// 		user := &domain.User{
// 			ID:       uuid.New(),
// 			Username: "new_test@example.com",
// 			Password: util.Sha256("password123"),
// 		}

// 		mockRepo.On("ExistsByUsername", signinRequest.Username).Return(false)
// 		mockRepo.On("PasswordCheck", signinRequest.Username, util.Sha256(signinRequest.Password)).Return(user.ID)
// 		mockRepo.On("Retrieve", signinRequest.Username, util.Sha256(signinRequest.Password)).Return(user)

// 		token, err := userService.Signin(signinRequest)

// 		assert.Error(t, err)
// 		assert.Empty(t, token)
// 		assert.Equal(t, "user with this username doesn't exist", err.Error())
// 	})
// }
