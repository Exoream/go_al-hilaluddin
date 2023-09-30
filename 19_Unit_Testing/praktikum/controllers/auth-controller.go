package controllers

import (
	"db_API/helpers"
	"db_API/models"
	"db_API/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

// check login request
func LoginController(c echo.Context) error {
	var loginUser = models.UserResponse{}
	c.Bind(&loginUser)
	
	data, token, err := repositories.QueryCheckLoginUser(loginUser.Email, loginUser.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("error login", "failed"))
	}

	response := helpers.SuccesResponseAuth(data.ID, data.Name, token)

	return c.JSON(http.StatusOK, helpers.SuccesResponse("login successful", response))
}
