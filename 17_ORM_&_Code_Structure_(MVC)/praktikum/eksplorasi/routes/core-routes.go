package routes

import (
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	UserRoutes(e)
	BlogRoutes(e)
	return e
}
