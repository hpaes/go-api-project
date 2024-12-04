package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hpaes/go-api-project/src/core/application/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAccountController_Signup(t *testing.T) {
	mockSignUp := NewMockSignup()
	mockGetAccount := NewMockGetAccount()
	mockLogger := NewMockLogger()

	controller := NewSignupController(mockSignUp, mockGetAccount, mockLogger)

	t.Run("should return 201 when signup is successful", func(t *testing.T) {
		input := usecase.SignupInput{
			Name:        "John Doe",
			Email:       "johnDoe@email.com",
			Cpf:         "123.456.789-09",
			CarPlate:    "ABC-1B34",
			IsPassenger: true,
			IsDriver:    false,
		}
		output := &usecase.SignupOutput{
			AccountId: "123",
		}

		mockSignUp.On("Execute", mock.Anything, input).Return(output, nil)
		mockLogger.On("LogInformation", "AccountController.Signup", mock.Anything).Once()

		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
		rr := httptest.NewRecorder()

		controller.Signup(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockSignUp.AssertExpectations(t)

	})

	t.Run("should return 400 when request method is not POST", func(t *testing.T) {
		mockLogger.On("LogInformation", "AccountController.Signup", mock.Anything).Once()
		mockLogger.On("LogError", "Error occurred", mock.Anything).Once()

		req, _ := http.NewRequest(http.MethodGet, "/signup", nil)
		rr := httptest.NewRecorder()

		controller.Signup(rr, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should return 400 when request payload is invalid", func(t *testing.T) {
		mockLogger.On("LogInformation", "AccountController.Signup", mock.Anything).Once()
		mockLogger.On("LogError", "Error occurred", mock.Anything).Once()
		req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer([]byte("invalid")))
		rr := httptest.NewRecorder()

		controller.Signup(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockLogger.AssertExpectations(t)
	})
}

func TestAccountController_GetAccount(t *testing.T) {
	mockSignUp := NewMockSignup()
	mockGetAccount := NewMockGetAccount()
	mockLogger := NewMockLogger()

	controller := NewSignupController(mockSignUp, mockGetAccount, mockLogger)

	t.Run("should return 200 when get account is successful", func(t *testing.T) {
		mockSignUp = NewMockSignup()
		mockGetAccount = NewMockGetAccount()
		mockLogger = NewMockLogger()
		controller = NewSignupController(mockSignUp, mockGetAccount, mockLogger)

		accountId := "123"
		output := &usecase.GetAccountOutput{
			AccountId:   "123",
			Name:        "John Doe",
			Email:       "johnDoe@email.com",
			Cpf:         "123.456.789-09",
			CarPlate:    "ABC-1B34",
			IsPassenger: true,
			IsDriver:    false}

		mockGetAccount.On("Execute", mock.Anything, accountId).Return(output, nil)
		mockLogger.On("LogInformation", "AccountController.GetAccount", mock.Anything).Once()

		req, _ := http.NewRequest(http.MethodGet, "/account?account_id=123", nil)
		rr := httptest.NewRecorder()

		controller.GetAccount(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockGetAccount.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should return 405 when request method is not GET", func(t *testing.T) {
		mockSignUp = NewMockSignup()
		mockGetAccount = NewMockGetAccount()
		mockLogger = NewMockLogger()
		controller = NewSignupController(mockSignUp, mockGetAccount, mockLogger)

		mockLogger.On("LogInformation", "AccountController.GetAccount", mock.Anything).Once()
		mockLogger.On("LogError", "Error occurred", mock.Anything).Once()

		req, _ := http.NewRequest(http.MethodPost, "/account", nil)
		rr := httptest.NewRecorder()

		controller.GetAccount(rr, req)

		assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should return 404 when account is not found", func(t *testing.T) {
		mockSignUp = NewMockSignup()
		mockGetAccount = NewMockGetAccount()
		mockLogger = NewMockLogger()
		controller = NewSignupController(mockSignUp, mockGetAccount, mockLogger)

		accountId := "123"

		mockGetAccount.On("Execute", mock.Anything, accountId).Return((*usecase.GetAccountOutput)(nil), nil)
		mockLogger.On("LogInformation", "AccountController.GetAccount", mock.Anything).Once()
		mockLogger.On("LogError", "Error occurred", mock.Anything).Once()

		req, _ := http.NewRequest(http.MethodGet, "/account?account_id=123", nil)
		rr := httptest.NewRecorder()

		controller.GetAccount(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		mockGetAccount.AssertExpectations(t)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should return 400 when account_id is missing", func(t *testing.T) {
		mockSignUp = NewMockSignup()
		mockGetAccount = NewMockGetAccount()
		mockLogger = NewMockLogger()
		controller = NewSignupController(mockSignUp, mockGetAccount, mockLogger)

		mockLogger.On("LogInformation", "AccountController.GetAccount", mock.Anything).Once()
		mockLogger.On("LogError", "Error occurred", mock.Anything).Once()

		req, _ := http.NewRequest(http.MethodGet, "/account", nil)
		rr := httptest.NewRecorder()

		controller.GetAccount(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		mockLogger.AssertExpectations(t)
	})

}
