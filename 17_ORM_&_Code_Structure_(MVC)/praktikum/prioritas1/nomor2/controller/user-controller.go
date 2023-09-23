package controller

import (
	"net/http"
	"strconv"

	"db_API/prioritas1/nomor2/helper"
	"db_API/prioritas1/nomor2/model"
	"db_API/prioritas1/nomor2/repository"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	users, err := repository.QuerySelectAllUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("success get all users", users))
}

// get user by id
func GetUserControllerById(c echo.Context) error {
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error convert", "failed" ))
	}

	user, err := repository.QuerySelectUserById(idUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found", "failed" ))
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("user found", user))	
}

// create new user
func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := repository.QueryCreateUser(&user); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

	return c.JSON(http.StatusOK, helper.SuccesResponse("success create new user", user))
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error convert", "failed"))
	}

	if err := repository.QueryDeleteUserById(idUser); err != nil {
		return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found", "failed"))
	}

	return c.JSON(http.StatusOK, helper.SuccesResponse("user deleted", "succes"))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	idUser, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error convert", "failed"))
	}

	updatedUser := new(model.User)
	if err := c.Bind(updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, helper.ErrorResponse("error binding data", "failed"))
	}

	if err := repository.QueryUpdateUserById(idUser, updatedUser); err != nil {
        return c.JSON(http.StatusNotFound, helper.ErrorResponse("user not found", "failed"))
    }

	return c.JSON(http.StatusOK, helper.SuccesResponse("succes user updated", updatedUser))
}
