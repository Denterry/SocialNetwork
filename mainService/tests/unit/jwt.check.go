package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Denterry/SocialNetwork/authServiceModernVersion/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Signup(request *model.SignupRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("successful signup", func(t *testing.T) {
		mockUserService := new(MockUserService)
		controller := userController{service: mockUserService}

		mockUserService.On("Signup", mock.Anything).Return(nil)

		router := gin.Default()
		router.POST("/signup", controller.Signup)

		signupRequest := model.SignupRequest{
			Username:             "test@example.com",
			Password:             "password123",
			PasswordConfirmation: "password123",
		}
		body, _ := json.Marshal(signupRequest)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"You have successfully registered!"}`, w.Body.String())

		mockUserService.AssertExpectations(t)
	})

	t.Run("validation failed", func(t *testing.T) {
		mockUserService := new(MockUserService)
		controller := userController{service: mockUserService}

		router := gin.Default()
		router.POST("/signup", controller.Signup)

		// Invalid signup request (missing password confirmation)
		signupRequest := model.SignupRequest{
			Username: "test@example.com",
			Password: "password123",
		}
		body, _ := json.Marshal(signupRequest)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.JSONEq(t, `{"message":"Validation failed"}`, w.Body.String())
	})

	t.Run("signup service error", func(t *testing.T) {
		mockUserService := new(MockUserService)
		controller := userController{service: mockUserService}

		mockUserService.On("Signup", mock.Anything).Return(assert.AnError)

		router := gin.Default()
		router.POST("/signup", controller.Signup)

		signupRequest := model.SignupRequest{
			Username:             "test@example.com",
			Password:             "password123",
			PasswordConfirmation: "password123",
		}
		body, _ := json.Marshal(signupRequest)

		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotAcceptable, w.Code)
		assert.JSONEq(t, `{"message":"assert.AnError general error for testing"}`, w.Body.String())

		mockUserService.AssertExpectations(t)
	})
}
