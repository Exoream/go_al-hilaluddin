package routes

import (
	"db_API/prioritas2/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {
	// Route Users
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserControllerById)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
}
