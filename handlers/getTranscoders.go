package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @description Get all the transcoders
// @host localhost:51000
// @Accept */*
// @Produce application/json
// @Param input_type query string false "input_type"
// @Param output_type query string false "output_type"
// @Success 200 {object} []Transcoder "OK"
// @Failure 500 {object} string "Internal Server Error"
// @Failure 400 {object} string "Invalid limit or skip."
// @Router /transcoders [get]
func (h *TranscoderHandler) GetTranscoders(c echo.Context) error {

	// Check if the request has a query parameter
	output_type := c.QueryParam("output_type")
	input_type := c.QueryParam("input_type")
	limit := c.QueryParam("limit")
	skip := c.QueryParam("skip")

	// If the request has a query parameter
	filter := bson.M{}

	// If output_type is present in the query parameter
	if output_type != "" {
		filter["output_type"] = output_type
	}

	// If input_type is present in the query parameter
	if input_type != "" {
		filter["input_type"] = input_type
	}

	// Array to hold the response
	var transcoders []Transcoder

	// Options to limit the number of documents returned
	opts := options.Find()
	if limit == "" {
		opts.SetLimit(10)
	} else {
		parsedLimit, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid limit.")
		}
		opts.SetLimit(parsedLimit)
	}

	// Options to skip the number of documents returned
	if skip == "" {
		opts.SetSkip(0)
	} else {
		parsedSkip, err := strconv.ParseInt(skip, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid skip.")
		}
		opts.SetSkip(parsedSkip)
	}

	// Getting the data from the database
	if data, err := h.Col.Find(context.Background(), filter, opts); err != nil {
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	} else {
		data.All(context.Background(), &transcoders)
		return c.JSON(http.StatusOK, transcoders)
	}
}
