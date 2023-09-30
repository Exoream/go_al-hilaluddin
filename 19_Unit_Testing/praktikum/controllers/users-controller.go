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

// get all users
func GetUsersController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	users, err := repositories.QuerySelectAllUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, helpers.ErrorResponse("failed to load data", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success get all users", users))
}

// get user by id
func GetUserControllerById(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}
	
	user, err := repositories.QuerySelectUserById(idUser)
	if err != nil {
		return c.JSON(http.StatusOK, helpers.ErrorResponse("user not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("user found", user))
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	result, err := repositories.QueryCreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed to create data", "failed"))
	}
	return c.JSON(http.StatusCreated, helpers.SuccesResponses("success", "success create new user", result))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	if err := repositories.QueryDeleteUserById(idUser); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("user deleted", "succes"))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	middlewares.ExtractTokenUserId(c)
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	updatedUser := new(models.User)
	c.Bind(updatedUser)
	updatedUser, err := repositories.QueryUpdateUserById(idUser, updatedUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("succes user updated", updatedUser))
}
