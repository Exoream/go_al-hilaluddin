package controllers

import (
	"bytes"
	"clean/api/entity"
	"clean/api/middlewares"
	"clean/api/models"
	"encoding/json"
	"errors"
	"fmt"

	mocks "clean/api/mocks"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsersCreateUser(t *testing.T) {
	mockData := models.UserResponse{
		Email:    "hilal@gmail.com",
		Password: "123",
	}

	request, err := json.Marshal(mockData)
	if err != nil {
		t.Error(t, err, "error")
	}

	e := echo.New()
	usecase := new(mocks.UseCaseInterface)

	t.Run("Success Create User", func(t *testing.T) {
		usecase.On("Create", mock.Anything).Return(1, nil).Once()

		controller := NewUserControllers(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(request))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := models.ResponseMessage{}

		if assert.NoError(t, controller.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			// fmt.Println("res", responseBody)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "success create new user", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Failed Create User", func(t *testing.T) {
		usecase.On("Create", mock.Anything).Return(0, errors.New("failed to insert data")).Once()
		controller := NewUserControllers(usecase)

		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(request))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := models.ResponseMessage{}

		if assert.NoError(t, controller.CreateUser(echoContext)) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed to create data", responseData.Message)
		}
		usecase.AssertExpectations(t)
	})
}

func TestLoginUser(t *testing.T) {
	t.Run("Success Login User", func(t *testing.T) {
		mockUseCase := new(mocks.UseCaseInterface)
		controller := NewUserControllers(mockUseCase)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/login", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockUseCase.On("CheckLogin", mock.Anything, mock.Anything).Return(entity.Main{}, "token", nil)

		err := controller.LoginUser(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		mockUseCase.AssertCalled(t, "CheckLogin", mock.Anything, mock.Anything)
	})

	t.Run("Failed Login User", func(t *testing.T) {
		usecase := new(mocks.UseCaseInterface)
		controller := NewUserControllers(usecase)

		e := echo.New()
		validLoginData := map[string]interface{}{
			"Email":    "validemail@example.com",
			"Password": "validpassword",
		}
		reqBody, _ := json.Marshal(validLoginData)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/login")

		expectedErrorMessage := "login failed"
		usecase.On("CheckLogin", validLoginData["Email"].(string), validLoginData["Password"].(string)).Return(entity.Main{}, "", errors.New(expectedErrorMessage)).Once()

		if assert.NoError(t, controller.LoginUser(c)) {
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
			responseBody := rec.Body.String()
			var responseData models.ResponseMessage
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if assert.NoError(t, err) {
				assert.Equal(t, expectedErrorMessage, responseData.Message)
			}
		}
		usecase.AssertExpectations(t)
	})
}

func TestGetUsersController(t *testing.T) {
	e := echo.New()
	usecase := new(mocks.UseCaseInterface)
	returnData := []entity.Main{
		{
			ID: 1, 
			Email: "aa@gmail.com", 
			Password: "123",
		},
	}

	t.Run("Success Get Users", func(t *testing.T) {
		usecase.On("FindAllUsers").Return(returnData, nil).Once()
		token, errToken := middlewares.CreateToken(returnData[0].ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserControllers(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		responseData := models.UsersResponseSuccess{}
		callFunc := middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllUser))(echoContext)

		if assert.NoError(t, callFunc) {
			responseBody := rec.Body.String()
			err := json.Unmarshal([]byte(responseBody), &responseData)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, returnData[0].Email, responseData.Data[0].Email)
		}
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get Users", func(t *testing.T) {
		usecase.On("FindAllUsers").Return(nil, errors.New("error get data")).Once()
		token, errToken := middlewares.CreateToken(returnData[0].ID)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserControllers(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllUser))(echoContext)
		responseBody := rec.Body.String()
		var responseData models.ResponseMessage
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "error get data", responseData.Message)
		usecase.AssertExpectations(t)
	})

	t.Run("Error Get Users No Token", func(t *testing.T) {
		token, errToken := middlewares.CreateToken(0)
		if errToken != nil {
			assert.Error(t, errToken)
		}
		srv := NewUserControllers(usecase)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		echoContext := e.NewContext(req, rec)
		echoContext.SetPath("/users")

		middlewares.JWTMiddleware()(echo.HandlerFunc(srv.GetAllUser))(echoContext)
		responseBody := rec.Body.String()
		var responseData models.ResponseMessage
		err := json.Unmarshal([]byte(responseBody), &responseData)
		fmt.Println("res", responseData)
		if err != nil {
			assert.Error(t, err, "error")
		}
		assert.Equal(t, "unauthorized", responseData.Message)
		usecase.AssertExpectations(t)
	})
}

