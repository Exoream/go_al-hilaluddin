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
func GetUsersController(c echo.Context) error {
	users, err := repositories.QuerySelectAllUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success get all users", users))
}

// get user by id
func GetUserControllerById(c echo.Context) error {
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	user, err := repositories.QuerySelectUserById(idUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ErrorResponse("user not found", "failed"))
	}

	return c.JSON(http.StatusOK, helpers.SuccesResponse("user found", user))	
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := repositories.QueryCreateUser(&user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusOK, helpers.SuccesResponse("success create new user", user))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
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
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error convert", "failed"))
	}

	updatedUser := new(models.User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("error binding data", "failed"))
	}

	updatedUser, err := repositories.QueryUpdateUserById(idUser, updatedUser)
    if err != nil {
        return c.JSON(http.StatusNotFound, helpers.ErrorResponse("error updating user", "failed"))
    }

	return c.JSON(http.StatusOK, helpers.SuccesResponse("succes user updated", updatedUser))
}

