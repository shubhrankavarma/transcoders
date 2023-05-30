package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Readyz is a readiness check API
// @description Readiness check API
// @Accept */*
// @Produce application/json
// @Success 200 {object} string "OK"
// @Failure 503 {object} string "Service Unavailable"
// @Router /readyz [get]
func (h *TranscoderHandler) Readyz(c echo.Context) error {
	if h.IsReady == nil || !h.IsReady.Load().(bool) {
		log.Errorf("Readyz API called and responded with status code 503")
		return c.JSON(http.StatusServiceUnavailable, nil)
	}
	log.Infof("Readyz API called and responded with OK")
	return c.JSON(http.StatusOK, "Welcome to Transcode Command Service")
}
