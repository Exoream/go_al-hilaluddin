package controllers

import (
	"net/http"
	"strconv"

	"db_API/eksplorasi/helpers"
	"db_API/eksplorasi/models"
	"db_API/eksplorasi/repositories"

	"github.com/labstack/echo/v4"
)

// get all users
func GetBlogController(c echo.Context) error {
	blogs, err := repositories.QuerySelectAllBlog()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success get all blogs", blogs))
}

// get user by id
func GetBlogControllerById(c echo.Context) error {
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

// create new user
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := repositories.QueryCreateBlog(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success create new blog", blog))
}

// delete user by id
func DeleteBlogController(c echo.Context) error {
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

// update user by id
func UpdateBlogController(c echo.Context) error {
	id := c.Param("id")
	idBlog, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	updatedBlog := new(models.Blog)
	if err := c.Bind(updatedBlog); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error binding data", "failed"))
	}

	updatedBlog, err := repositories.QueryUpdateBlogById(idBlog, updatedBlog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("error updating blog", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("succes blog updated", updatedBlog))
}
