package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

type response struct {
	Value string `example:"Transcoder added successfully."`
}

// AddTranscoder is a post transcoder API
// @title Add Transcoder API
// @description Adds the transcoder to the database
// @host localhost:51000
// @BasePath /
// @Router /transcoder [post]
// @schemes http
// @Param Authorization header string true "JWT Token"
// @Param input body Transcoder true "Transcoder" required=true
// @Accept application/json
// @Produce application/json
// @Success 201 {object} response "Transcoder added successfully."
// @Failure 400 {object} string "Invalid request payload." example:"Invalid request payload."
// @Failure 409 {object} string "Transcoder with the same output type and input type already exists." example:"Transcoder with the same output type and input type already exists."
// @Failure 422 {object} string "Unable to pass the request payload." example:"Unable to pass the request payload."
// @Failure 500 {object} string "Unable to process the request." example:"Unable to process the request."
func (h *TranscoderHandler) AddTranscoder(c echo.Context) error {

	// Variable to hold the request payload
	var trancoder Transcoder

	// Binding the request payload to the variable
	c.Echo().Validator = &TranscoderValidator{validator: v}
	if err := c.Bind(&trancoder); err != nil {
		log.Errorf("Error while binding the request: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, "Unable to pass the request payload.")
	}

	// Validating the request payload
	if err := c.Validate(trancoder); err != nil {
		log.Errorf("Error while validating the request: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid request payload.")
	}

	// Check if any transcoder with the same output type and input type exists
	filter := bson.M{
		"output_type": trancoder.OutputType,
		"input_type":  trancoder.InputType,
	}

	// Checking if the transcoder already exists
	res := h.Col.FindOne(context.Background(), filter)
	if res.Err() == nil {
		return c.JSON(http.StatusConflict, "Transcoder with the same output type and input type already exists.")
	}

	// Add created_at and updated_at
	trancoder.CreatedAt = time.Now()
	trancoder.UpdatedAt = time.Now()

	// Inserting the request payload to the database
	if _, err := h.Col.InsertOne(context.TODO(), trancoder); err != nil {
		log.Errorf("Error while inserting the request: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	}

	// Returning the response
	return c.JSON(http.StatusCreated, "Transcoder added successfully.")
}
