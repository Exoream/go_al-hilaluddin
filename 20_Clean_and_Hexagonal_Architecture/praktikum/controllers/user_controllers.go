package controllers

import (
	"clean/api/entity"
	"clean/api/helpers"
	"clean/api/middlewares"
	"clean/api/models"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUseCase entity.UseCaseInterface
}

func NewUserControllers(uc entity.UseCaseInterface) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (uco *UserController) GetAllUser(c echo.Context) error {
	idToken := middlewares.ExtractToken(c)
	if idToken < 1 {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "unauthorized",
		})
	}
	responseData, err := uco.userUseCase.FindAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ErrorResponse("error get data"))
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("Success receive user data", responseData))
}

func (uco *UserController) CreateUser(c echo.Context) error {
	user := new(models.UserResponse)
	c.Bind(&user)

	data := entity.Main{
		Email:    user.Email,
		Password: user.Password,
	}

	_, errCreate := uco.userUseCase.Create(data)
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed to create data"))
	}

	return c.JSON(http.StatusCreated, helpers.SuccessWithDataResponse("success create new user", user))
}

func (s *UserController) LoginUser(c echo.Context) error {
	var login models.UserLoginResponse
	c.Bind(&login)

	user, token, err := s.userUseCase.CheckLogin(login.Email, login.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ErrorResponse("login failed"))
	}

	response := models.UserLoginResponse{
		Email:    user.Email,
		Password: "",
		Token:    token,
	}

	return c.JSON(http.StatusOK, helpers.SuccessWithDataResponse("login successful", response))
}
