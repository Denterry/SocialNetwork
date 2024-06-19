package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Denterry/SocialNetwork/mainService/internal/controller"
	mock_service "github.com/Denterry/SocialNetwork/mainService/internal/service/mocks"
	"github.com/Denterry/SocialNetwork/mainService/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, request *model.SignupRequest)

	testTable := []struct {
		name                string
		inputBody           string
		inputUser           *model.SignupRequest
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody gin.H
	}{
		{
			name:      "Ok",
			inputBody: `{"username": "test@example.com", "password": "qwerty", "password_confirmation": "qwerty"}`,
			inputUser: &model.SignupRequest{
				Username:             "test@example.com",
				Password:             "qwerty",
				PasswordConfirmation: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockUserService, request *model.SignupRequest) {
				s.EXPECT().Signup(request).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: gin.H{"message": "You have successfully registered!"},
		},
		{
			name:                "Empty fields",
			inputBody:           `{"password": "qwerty", "password_confirmation": "qwerty"}`,
			mockBehavior:        func(s *mock_service.MockUserService, request *model.SignupRequest) {},
			expectedStatusCode:  400,
			expectedRequestBody: gin.H{"message": "Validation failed"},
		},
		{
			name:      "Passwords don't match (service failure)",
			inputBody: `{"username": "test@example.com", "password": "qwerty", "password_confirmation": "qwrty"}`,
			inputUser: &model.SignupRequest{
				Username:             "test@example.com",
				Password:             "qwerty",
				PasswordConfirmation: "qwrty",
			},
			mockBehavior: func(s *mock_service.MockUserService, request *model.SignupRequest) {
				s.EXPECT().Signup(request).Return(errors.New("password and confirm password must match"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: gin.H{"message": "password and confirm password must match"},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserService(c)
			testCase.mockBehavior(user, testCase.inputUser)

			stat := mock_service.NewMockStatisticsServiceClient(c)

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()

			controller := controller.NewUserController(router, user, nil, stat)

			// Test server
			router.POST("/sign-up", controller.Signup)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			var responseBody gin.H
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedRequestBody, responseBody)
		})
	}
}

func TestHandler_changeInfo(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUserService, request *model.ChangeInfoRequest)

	testTable := []struct {
		name                string
		inputBody           string
		inputInfo           *model.ChangeInfoRequest
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody gin.H
	}{
		{
			name: "Ok",
			inputBody: `{"username": "test@example.com", 
						"name": "name",
						"surname": "surname",
						"email": "den@gmail.com"}`,
			inputInfo: &model.ChangeInfoRequest{
				Username: "test@example.com",
				Name:     "name",
				Surname:  "surname",
				Email:    "den@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockUserService, request *model.ChangeInfoRequest) {
				s.EXPECT().ChangeInfo(request).Return(nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: gin.H{"message": "You have successfully changed your personal info!"},
		},
		{
			name: "Empty fields",
			inputBody: `{"username": "test@example.com", 
						"name": "name"}`,
			mockBehavior:        func(s *mock_service.MockUserService, request *model.ChangeInfoRequest) {},
			expectedStatusCode:  400,
			expectedRequestBody: gin.H{"message": "Validation failed"},
		},
		{
			name: "Service failure",
			inputBody: `{"username": "test@example.com", 
						"name": "name",
						"surname": "surname",
						"email": "den@gmail.com"}`,
			inputInfo: &model.ChangeInfoRequest{
				Username: "test@example.com",
				Name:     "name",
				Surname:  "surname",
				Email:    "den@gmail.com",
			},
			mockBehavior: func(s *mock_service.MockUserService, request *model.ChangeInfoRequest) {
				s.EXPECT().ChangeInfo(request).Return(errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: gin.H{"message": "service failure"},
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUserService(c)
			testCase.mockBehavior(user, testCase.inputInfo)

			stat := mock_service.NewMockStatisticsServiceClient(c)

			gin.SetMode(gin.ReleaseMode)
			router := gin.New()

			controller := controller.NewUserController(router, user, nil, stat)

			// Test server
			router.PUT("/change-info", controller.ChangeInfo)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/change-info", bytes.NewBufferString(testCase.inputBody))

			// Perform request
			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			var responseBody gin.H
			err := json.Unmarshal(w.Body.Bytes(), &responseBody)
			assert.NoError(t, err)
			assert.Equal(t, testCase.expectedRequestBody, responseBody)
		})
	}
}
