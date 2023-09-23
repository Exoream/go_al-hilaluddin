package routes

import (
	"db_API/prioritas2/controllers"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	// Routes Books
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookControllerById)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
