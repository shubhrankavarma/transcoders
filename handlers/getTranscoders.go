package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

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

	// Check if the request has a query parameter
	outputType := c.QueryParam("output_type")
	inputType := c.QueryParam("input_type")
	codec := c.QueryParam("codec")
	descriptor := c.QueryParam("descriptor")
	pageSizeQueryParam := c.QueryParam("page_size")
	pageQueryParam := c.QueryParam("page")
	var err error

	// If the request has a query parameter
	filter := bson.M{}

	// If output_type is present in the query parameter
	if outputType != "" {
		filter["output_type"] = outputType
	}

	// If input_type is present in the query parameter
	if inputType != "" {
		filter["input_type"] = inputType
	}

	// If codec is present in the query parameter
	if codec != "" {
		filter["codec"] = codec
	}

	// If descriptor is present in the query parameter
	if descriptor != "" {
		filter["descriptor"] = descriptor
	}

	// If page is not present in the query parameter
	page := 1
	if pageQueryParam != "" {
		page, err = strconv.Atoi(pageQueryParam)
		log.Infof("page %v", page)
		if err != nil {
			log.Errorf("Unable to parse the query param page: %v", err)
			return c.JSON(http.StatusBadRequest, "Unable to parse the query param page.")
		}
	}

	// If page_size is not present in the query parameter
	limit := 10
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
