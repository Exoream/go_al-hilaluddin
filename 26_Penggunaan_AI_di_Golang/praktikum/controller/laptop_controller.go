package controller

import (
	"fmt"
	"laptop/entity"
	"laptop/respon"
	"laptop/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LaptopController struct {
	LaptopUsecase usecase.LaptopUsecase
}

func NewLaptopController(usecase usecase.LaptopUsecase) *LaptopController {
	return &LaptopController{
		LaptopUsecase: usecase,
	}
}

func (h *LaptopController) RecommendLaptop(c echo.Context) error {
	var requestData respon.RecommendRequest
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	recommendation := fmt.Sprintf("Here's a recommendation for a %s laptop brand with %s RAM, %s CPU, a %s screen size, and a budget of Rp %.0f, suitable for %s.", requestData.Brand, requestData.RAM, requestData.CPU, requestData.ScreenSize, requestData.Budget, requestData.Purpose)

	answer, err := h.LaptopUsecase.RecommendLaptop(entity.LaptopParams{
		UserInput:  recommendation,

	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error in generating recommendation")
	}

	responseData := respon.RecommendResponse{
		Status: "success",
		Data:   answer,
	}

	return c.JSON(http.StatusOK, responseData)
}

