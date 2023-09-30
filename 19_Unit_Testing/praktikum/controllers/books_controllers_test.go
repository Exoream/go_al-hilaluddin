package controllers

import (
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

var (
	book = models.Book{
		Judul:    "cakrawala",
		Penulis:  "budi",
		Penerbit: "gramedia",
	}
)

type BooksResponseSuccess struct {
	Status  string
	Message string
	Data    []models.Book
}

type SingleBookResponseSuccess struct {
	Status  string
	Message string
	Data    models.Book
}

func InsertDataBook() error {
	var err error
	if err = config.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBooksController(t *testing.T) {
	t.Run("Get All Books", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBooksController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			type Response struct {
				Message string                `json:"message"`
				Data    []models.BookResponse `json:"data"`
			}
			var responseData Response
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "success get all books", responseData.Message)
			assert.Equal(t, 1, len(responseData.Data))
		}
	})

	t.Run("Failed Get All Books", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		config.DB.Migrator().DropTable(&models.Book{})

		req := httptest.NewRequest(http.MethodGet, "/books", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBooksController))(c)

		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			body := rec.Body.String()
			type Response struct {
				Message string                `json:"message"`
				Data    []models.BookResponse `json:"data"`
			}
			var book ResponseFailed
			err = json.Unmarshal([]byte(body), &book)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed get all books", book.Message)
			assert.Equal(t, "failed", book.Status)
		}
	})
}

func TestGetBookControllerById(t *testing.T) {
	t.Run("Get Book By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBookControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			var responseData SingleBookResponseSuccess
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "book found", responseData.Message)
			assert.Equal(t, "cakrawala", responseData.Data.Judul)
		}
	})

	t.Run("Not Found Book By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBookControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBookResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "book not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/books/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBookControllerById))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})
}

func TestCreateBookController(t *testing.T) {
	t.Run("Pembuatan book berhasil", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()

		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/books", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(CreateBookController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			var response map[string]interface{}
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "success create new book", response["message"])
			assert.NotNil(t, response["data"])
		}
	})

	t.Run("Pembuatan book gagal", func(t *testing.T) {
		e := InitEchoTestAPI()

		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		config.DB.Migrator().DropTable(&models.Book{})

		req := httptest.NewRequest(http.MethodPost, "/books", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(CreateBookController))(c)

		if assert.NoError(t, cal) {
			var book ResponseFailed
			err := json.Unmarshal(rec.Body.Bytes(), &book)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", book.Status)
			assert.Equal(t, "failed to create data", book.Message)
		}
	})
}

func TestDeleteBookController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/books/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBookController))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Not Found Book By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/books/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBookController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBookResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "book not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Delete Book By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/books/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBookController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var responseData ResponseId
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "book deleted", responseData.Message)
		}
	})
}

func TestUpdateBookController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/books/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBookController))(c)
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
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/books/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBookController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBookResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "book not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Update Data", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBook()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/books/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBookController))(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var errorResponse ResponseFailed
			err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)

			assert.Equal(t, "succes book updated", errorResponse.Message)
		}
	})
}

// 	t.Run("Eror Binding Data", func(t *testing.T) {
// 		e := InitEchoTestAPI()
// 		token, err := middlewares.CreateToken(1)
// 		assert.NoError(t, err)

// 		// Membuat payload JSON yang tidak valid (contoh: data tidak lengkap)
// 		invalidJSON := `{"name": "John"}`

// 		req := httptest.NewRequest(http.MethodPut, "/users/1", strings.NewReader(invalidJSON))
// 		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
// 		c.SetParamNames("id")
// 		c.SetParamValues("1")

// 		// Memanggil fungsi UpdateUserController
// 		errorResponse := helpers.ErrorResponse("error binding data", "failed")

//     	if assert.NoError(t, middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateUserController))(c)) {
//         // Check the HTTP status code
//         assert.Equal(t, http.StatusBadRequest, rec.Code)

//         // Parse the response to verify the error message
//         var jsonResponse map[string]interface{}
//         err := json.Unmarshal(rec.Body.Bytes(), &jsonResponse)
//         assert.NoError(t, err)

//         // Assert the expected error message
//         assert.Equal(t, errorResponse, jsonResponse)
//     }
// })

// t.Run("Successful Update", func(t *testing.T) {
// 	e := InitEchoTestAPI()
// 	InsertDataUser()
// 	token, err := middlewares.CreateToken(1)
// 	assert.NoError(t, err)

//     updatedUserData := map[string]interface{}{
//         "Name": "Updated Name",
//     }

//     updatedUserDataJSON, err := json.Marshal(updatedUserData)
//     assert.NoError(t, err)

//     req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader(updatedUserDataJSON))
//     req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
//     rec := httptest.NewRecorder()
//     c := e.NewContext(req, rec)
//     c.SetParamNames("id")
//     c.SetParamValues("1")

//     if assert.NoError(t, middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateUserController))(c)) {
//         assert.Equal(t, http.StatusOK, rec.Code)

//         var responseData UsersResponseSuccess
//         err = json.Unmarshal(rec.Body.Bytes(), &responseData)
//         assert.NoError(t, err)

//         if len(responseData.Data) > 0 {
// 			assert.Equal(t, "Nama yang Diharapkan", responseData.Data[0].Name)
// 		} else {
// 			assert.Fail(t, "Tidak ada pengguna yang ditemukan dalam responseData.Data")
// 		}
//     }
// })
