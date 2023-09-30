package controllers

import (
	"bytes"
	"db_API/config"
	"db_API/middlewares"
	"db_API/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.InitDBTest()
	e := echo.New()
	return e
}

type UsersResponseSuccess struct {
	Status  string
	Message string
	Data    []models.User
}

type SingleUserResponseSuccess struct {
	Status  string
	Message string
	Data    models.User
}

type ResponseFailed struct {
	Status  string
	Message string
}

type ResponseId struct {
	Status  string
	Message string
}

var (
	user = models.User{
		Name:     "budi",
		Password: "123asd",
		Email:    "budi@gmail.com",
	}
)

func InsertDataUser() error {
	var err error
	if err = config.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}


func TestLogin(t *testing.T) {
	t.Run("Login Berhasil", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = LoginController(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "login successful", response["message"])
		assert.NotNil(t, response["data"])
	})

	t.Run("Error Login", func(t *testing.T) {
		e := InitEchoTestAPI()

		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err = LoginController(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var response ResponseFailed
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "error login", response.Message)
	})
}



func TestGetUserController(t *testing.T) {
	t.Run("Get All Users", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetUsersController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			type Response struct {
				Message string                `json:"message"`
				Data    []models.UserResponse `json:"data"`
			}
			var responseData Response
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "success get all users", responseData.Message)
			assert.Equal(t, 1, len(responseData.Data))
		}
	})
}

func TestGetUserControllerById(t *testing.T) {
	t.Run("Get User By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetUserControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			var responseData SingleUserResponseSuccess
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "user found", responseData.Message)
			assert.Equal(t, "budi", responseData.Data.Name)
		}
	})

	t.Run("Not Found User By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetUserControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleUserResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "user not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetUserControllerById))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})
}

func TestCreateUserController(t *testing.T) {
	t.Run("Pembuatan pengguna berhasil", func(t *testing.T) {
		e := InitEchoTestAPI()

		// Mengubah data user menjadi JSON
		requestBody, err := json.Marshal(user)
		assert.NoError(t, err)

		// Membuat HTTP request dengan method POST dan data JSON
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(requestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Memanggil handler CreateUserController
		err = CreateUserController(c)

		// Memeriksa kode status HTTP (dalam kasus ini, seharusnya 201 Created)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		// Memeriksa respons JSON yang dikembalikan
		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "success create new user", response["message"])
		assert.NotNil(t, response["data"])
	})

	t.Run("Pembuatan pengguna gagal", func(t *testing.T) {
		e := InitEchoTestAPI()

		config.DB.Migrator().DropTable(&models.User{})

		req := httptest.NewRequest(http.MethodPost, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, CreateUserController(c)) {
			var user ResponseFailed
			err := json.Unmarshal(rec.Body.Bytes(), &user)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", user.Status)
		}
	})
}

func TestDeleteUserController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/users/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteUserController))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Not Found User By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteUserController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleUserResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "user not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Delete User By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/users/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteUserController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var responseData ResponseId
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "user deleted", responseData.Message)
		}
	})
}

func TestUpdateUserController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/users/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateUserController))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Not Found User By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/users/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateUserController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleUserResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "user not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Update Data", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataUser()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/users/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateUserController))(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var errorResponse ResponseFailed
			err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)

			assert.Equal(t, "succes user updated", errorResponse.Message)
		}
	})
}
