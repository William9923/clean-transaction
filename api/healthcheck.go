package api

import (
	"net/http"

	"github.com/William9923/clean-transaction/api/dto"
	"github.com/labstack/echo/v4"
)

func HealthCheck(e *echo.Echo) {
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.ResponseEntity{
			Code:    "success",
			Message: "Healthy",
		})
	})
}
