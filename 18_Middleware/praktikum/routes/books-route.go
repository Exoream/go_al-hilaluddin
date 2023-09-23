package routes

import (
	"db_API/controllers"
	"db_API/middlewares"

	"github.com/labstack/echo/v4"
)

func BookRoutes(e *echo.Echo) {
	// Routes Books

	bookAunthen := e.Group("/books", middlewares.JWTMiddleware())
	bookAunthen.GET("", controllers.GetBooksController)
	bookAunthen.GET("/:id", controllers.GetBookControllerById)
	bookAunthen.POST("", controllers.CreateBookController)
	bookAunthen.DELETE("/:id", controllers.DeleteBookController)
	bookAunthen.PUT("/:id", controllers.UpdateBookController)
}
