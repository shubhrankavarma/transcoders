package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthz is a health check API
// @description Health check API
// @host localhost:52000
// @Accept */*
// @Produce application/json
// @Success 200 {object} string "OK"
// @Failure 500 {object} string "Internal Server Error"
// @Router / [get]
func (h *Transcoder) Healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "Welcome to Platforms Service")
}
