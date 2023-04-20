package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// GetTranscoders is a handler function to get all the transcoders
// @description Get all the transcoders
// @host localhost:51000
// @Accept */*
// @Produce application/json
// @Param input_type query string false "input_type"
// @Param output_type query string false "output_type"
// @Success 200 {object} []Transcoder "OK"github.com/amagi/kafkaConsumer
// @Failure 500 {object} string "Internal Server Error"
// @Failure 400 {object} string "Invalid limit or skip."
// @Router /transcoders [get]
func (h *TranscoderHandler) GetTranscoders(c echo.Context) error {

	filter := make(map[string]interface{})
	for k, v := range c.QueryParams() {
		if k != "page" && k != "page_size" {
			filter[k] = v[0]
		}
	}
	var err error

	// Getting the page_size and page from the query parameter
	pageSizeQueryParam := c.QueryParam("page_size")
	pageQueryParam := c.QueryParam("page")

	page := h.Cfg.PageNo
	if page == 0 {
		page = 1
	}

	// If page is present in the query parameter
	if pageQueryParam != "" {
		page, err = strconv.Atoi(pageQueryParam)
		log.Infof("page %v", page)
		if err != nil {
			log.Errorf("Unable to parse the query param page: %v", err)
			return c.JSON(http.StatusBadRequest, "Unable to parse the query param page.")
		}
	}

	// All the transcoders are active
	filter["status"] = "active"

	limit := (h.Cfg.PageSize)
	if limit == 0 {
		limit = 15
	}
	// If page_size is present in the query parameter
	if pageSizeQueryParam != "" {
		limit, err = strconv.Atoi(pageSizeQueryParam)
		log.Infof("page_size %v", limit)
		if err != nil {
			log.Errorf("Unable to parse the query param page_size: %v", err)
			return c.JSON(http.StatusBadRequest, "Unable to parse the query param page_size.")
		}
	}

	// Array to hold the response
	var transcoders []Transcoder

	// Creating the options for the query
	opts := newMongoPaginate(limit, page).getPaginatedOpts()

	// Getting the data from the database
	if data, err := h.Col.Find(context.Background(), filter, opts); err != nil {
		log.Errorf("Unable to get the data from the database: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	} else {
		data.All(context.Background(), &transcoders)
		return c.JSON(http.StatusOK, transcoders)
	}
}
