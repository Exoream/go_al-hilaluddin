package routes

import (
	"db_API/controllers"
	"db_API/middlewares"

	"github.com/labstack/echo/v4"
)

func BlogRoutes(e *echo.Echo) {
	// Route Blogs
	blogAunthen := e.Group("/blogs", middlewares.JWTMiddleware())
	blogAunthen.GET("", controllers.GetBlogController)
	blogAunthen.GET("/:id", controllers.GetBlogControllerById)
	blogAunthen.POST("", controllers.CreateBlogController)
	blogAunthen.DELETE("/:id", controllers.DeleteBlogController)
	blogAunthen.PUT("/:id", controllers.UpdateBlogController)
}
