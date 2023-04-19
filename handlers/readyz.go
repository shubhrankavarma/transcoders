package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

//Readyz is a readiness check API
//@description Readiness check API
//@host localhost:8004
//@Accept */*

func (h *TranscoderHandler) Readyz(c echo.Context) error {
	if h.IsReady == nil || !h.IsReady.Load().(bool) {
		log.Errorf("Readyz API called and responded with status code 503")
		return c.JSON(http.StatusServiceUnavailable, nil)
	}
	log.Infof("Readyz API called and responded with OK")
	return c.JSON(http.StatusOK, "Welcome to Transcode Command Service")
}
