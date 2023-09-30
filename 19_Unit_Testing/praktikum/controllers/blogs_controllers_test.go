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
	blog = models.Blog{
		UserID:    1,
		Judul:  "nasi goreng",
		Konten: "food",
	}
)

type BlogsResponseSuccess struct {
	Status  string
	Message string
	Data    []models.Book
}

type SingleBlogResponseSuccess struct {
	Status  string
	Message string
	Data    models.Book
}

func InsertDataBlog() error {
	var err error
	if err = config.DB.Create(&blog).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBlogController(t *testing.T) {
	t.Run("Get All Blogs", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/blog", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBlogController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			type Response struct {
				Message string                `json:"message"`
				Data    []models.BlogResponse `json:"data"`
			}
			var responseData Response
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "success get all blogs", responseData.Message)
			assert.Equal(t, 1, len(responseData.Data))
		}
	})

	t.Run("Failed Get All Blogs", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		config.DB.Migrator().DropTable(&models.Blog{})

		req := httptest.NewRequest(http.MethodGet, "/blogs", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBlogController))(c)

		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			body := rec.Body.String()
			type Response struct {
				Message string                `json:"message"`
				Data    []models.BookResponse `json:"data"`
			}
			var blog ResponseFailed
			err = json.Unmarshal([]byte(body), &blog)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed get all blogs", blog.Message)
			assert.Equal(t, "failed", blog.Status)
		}
	})
}

func TestGetBlogControllerById(t *testing.T) {
	t.Run("Get Blog By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/blogs/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBlogControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()
			var responseData SingleBlogResponseSuccess
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, "blog found", responseData.Message)
			assert.Equal(t, "nasi goreng", responseData.Data.Judul)
		}
	})

	t.Run("Not Found Blog By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/blogs/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBlogControllerById))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBlogResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "blog not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/blogs/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(GetBlogControllerById))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "error convert", response.Message)
		assert.Equal(t, "failed", response.Status)
		}
	})
}

func TestCreateBlogController(t *testing.T) {
	t.Run("Pembuatan blog berhasil", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()

		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)
		

		req := httptest.NewRequest(http.MethodPost, "/blogs", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(CreateBlogController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, rec.Code)

			var response map[string]interface{}
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "success create new blog", response["message"])
			assert.NotNil(t, response["data"])
		}
	})

	t.Run("Pembuatan blog gagal", func(t *testing.T) {
		e := InitEchoTestAPI()

		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)
	
		config.DB.Migrator().DropTable(&models.Blog{})

		req := httptest.NewRequest(http.MethodPost, "/blogs", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(CreateBlogController))(c)

		if assert.NoError(t, cal) {
			var blog ResponseFailed
			err := json.Unmarshal(rec.Body.Bytes(), &blog)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "failed", blog.Status)
			assert.Equal(t, "failed to create data", blog.Message)
		}
	})
}

func TestDeleteBlogController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/blogs/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBlogController))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
			err = json.Unmarshal(rec.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, "error convert", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Not Found Blog By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodGet, "/blogs/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBlogController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBlogResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "blog not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Delete Blog By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodDelete, "/blogs/:id", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(DeleteBlogController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var responseData ResponseId
			err = json.Unmarshal([]byte(body), &responseData)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, "blog deleted", responseData.Message)
		}
	})
}

func TestUpdateBlogController(t *testing.T) {
	t.Run("error convert", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/blogs/invalid", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBlogController))(c)
		if assert.NoError(t, cal) {
			var response ResponseId
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)

		assert.Equal(t, "error convert", response.Message)
		assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Not Found Blog By Id", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/blogs/:id", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("100")

		cal := middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBlogController))(c)
		if assert.NoError(t, cal) {
			assert.NoError(t, err)
			body := rec.Body.String()

			var response SingleBlogResponseSuccess
			body = rec.Body.String()
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "blog not found", response.Message)
			assert.Equal(t, "failed", response.Status)
		}
	})

	t.Run("Update Data", func(t *testing.T) {
		e := InitEchoTestAPI()
		InsertDataBlog()
		token, err := middlewares.CreateToken(1)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/blogs/1", nil)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("1")

		if assert.NoError(t, middlewares.JWTMiddleware()(echo.HandlerFunc(UpdateBlogController))(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var errorResponse ResponseFailed
			err = json.Unmarshal(rec.Body.Bytes(), &errorResponse)
			assert.NoError(t, err)

			assert.Equal(t, "succes blog updated", errorResponse.Message)
		}
	})
}