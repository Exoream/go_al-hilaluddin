package routes

import (
	"db_API/controllers"
	"db_API/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	// Route Users
	userNotAuthen := e.Group("/users")
	userNotAuthen.POST("", controllers.CreateUserController)
	userNotAuthen.POST("/login", controllers.LoginController)

	userAuthen := e.Group("/users", middlewares.JWTMiddleware())
	userAuthen.GET("", controllers.GetUsersController)
	userAuthen.GET("/:id", controllers.GetUserControllerById)
	userAuthen.DELETE("/:id", controllers.DeleteUserController)
	userAuthen.PUT("/:id", controllers.UpdateUserController)
}
