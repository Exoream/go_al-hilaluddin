package router

import (
	"clean/api/controllers"
	"clean/api/middlewares"
	"clean/api/repositories"
	"clean/api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controllers.NewUserControllers(userUsecase)

	AuthenUsers := e.Group("/users", middlewares.JWTMiddleware())
	AuthenUsers.GET("", userController.GetAllUser)

	notAuthenUsers := e.Group("/users")
	notAuthenUsers.POST("", userController.CreateUser)
	notAuthenUsers.POST("/login", userController.LoginUser)

}
