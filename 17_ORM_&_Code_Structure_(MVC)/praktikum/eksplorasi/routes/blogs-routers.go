package routes

import (
	"db_API/eksplorasi/controllers"

	"github.com/labstack/echo/v4"
)

func BlogRoutes(e *echo.Echo) {
	// Route Users
	e.GET("/blogs", controllers.GetBlogController)
	e.GET("/blogs/:id", controllers.GetBlogControllerById)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
}