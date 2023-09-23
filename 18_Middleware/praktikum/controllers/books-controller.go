package controllers

import (
	"net/http"
	"strconv"

	"db_API/helpers"
	"db_API/models"
	"db_API/repositories"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := repositories.QuerySelectAllBook()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success get all books", books))
}

// get book by id
func GetBookControllerById(c echo.Context) error {
	id := c.Param("id")
	idBook, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	book, err := repositories.QuerySelectBookById(idBook)
	if err != nil {
		return c.JSON(http.StatusNotFound,  helpers.ErrorResponse("book not found", "failed" ))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("book found", book))	
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := repositories.QueryCreateBook(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success create new book", book))
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	idBook, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	if err := repositories.QueryDeleteBookById(idBook); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("book not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("book deleted", "succes"))
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	idBook, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	updatedBook := new(models.Book)
	if err := c.Bind(updatedBook); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error binding data", "failed"))
	}

	updatedBook, err := repositories.QueryUpdateBookById(idBook, updatedBook)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("book not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("succes book updated", updatedBook))
}
