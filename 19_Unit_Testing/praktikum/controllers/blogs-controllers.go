package controllers

import (
	"net/http"
	"strconv"

	"db_API/helpers"
	"db_API/middlewares"
	"db_API/models"
	"db_API/repositories"

	"github.com/labstack/echo/v4"
)

func GetBlogController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	blogs, err := repositories.QuerySelectAllBlog()

	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed get all blogs", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success get all blogs", blogs))
}

func GetBlogControllerById(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idBlog, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	blog, err := repositories.QuerySelectBlogById(idBlog)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("blog not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("blog found", blog))
}


func CreateBlogController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	blog := models.Blog{}
	c.Bind(&blog)

	result, err := repositories.QueryCreateBlog(&blog)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed to create data", "failed"))
	}
	return c.JSON(http.StatusOK, helpers.SuccesResponse("success create new blog", result))
}


func DeleteBlogController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idBlog, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	if err := repositories.QueryDeleteBlogById(idBlog); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("blog not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("blog deleted", "succes"))
}

func UpdateBlogController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idBlog, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	updatedBlog := new(models.Blog)
	c.Bind(updatedBlog)
	updatedBlog, err := repositories.QueryUpdateBlogById(idBlog, updatedBlog)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("blog not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("succes blog updated", updatedBlog))
}
