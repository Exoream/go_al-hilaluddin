package router

import (
	"laptop/controller"
	"laptop/usecase"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func NewRouter() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}


	openaiKey := os.Getenv("OPENAI_API_KEY")
	laptopUsecase := usecase.NewUserService(nil, openaiKey)
	laptopController := controller.NewLaptopController(*laptopUsecase)

	// Definisikan route
	e.POST("/recommend-laptop", laptopController.RecommendLaptop)

	// Mulai server Echo
	e.Start(":8080")
}
