package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

// AddTranscoder is a post transcoder API
// @title Add Transcoder API
// @description Adds the transcoder to the database
// @Router /commands [post]
// @schemes http
// @Param Authorization header string true "JWT Token"
// @Param input body Transcoder true "Transcoder" required=true
// @Accept application/json
// @Produce application/json
// @Success 201 {object} string "Transcoder added successfully."
// @Failure 400 {object} string "Invalid request payload." example:"Invalid request payload."
// @Failure 409 {object} string "Transcoder with the same output type and input type already exists." example:"Transcoder with the same output type and input type already exists."
// @Failure 422 {object} string "Unable to pass the request payload." example:"Unable to pass the request payload."
// @Failure 500 {object} string "Unable to process the request." example:"Unable to process the request."
func (h *TranscoderHandler) AddTranscoder(c echo.Context) error {

	// Variable to hold the request payload
	var transcoder Transcoder

	// Binding the request payload to the variable
	c.Echo().Validator = &TranscoderValidator{validator: v}
	if err := c.Bind(&transcoder); err != nil {
		log.Errorf("Error while binding the request: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, "Unable to pass the request payload.")
	}

	// Validating the request payload
	if err := c.Validate(transcoder); err != nil {
		log.Errorf("Error while validating the request: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid request payload.")
	}

	// Check if any transcoder with the same output type, input type and codec exists
	filter := bson.M{
		"operation":  transcoder.Operation,
		"asset_type": transcoder.AssetType,
	}
	if transcoder.InputType != "" {
		filter["input_type"] = transcoder.InputType
	}
	if transcoder.OutputType != "" {
		filter["output_type"] = transcoder.OutputType
	}

	// Checking if the transcoder already exists
	res := h.Col.FindOne(context.Background(), filter)
	if res.Err() == nil {
		log.Error("Transcoder with the same parameters already exists.")
		return c.JSON(http.StatusConflict, "Transcoder with the same parameters already exists.")
	}

	// Add created_at and updated_at
	transcoder.CreatedAt = time.Now()
	transcoder.UpdatedAt = time.Now()

	// Set the status to active
	transcoder.Status = "active"

	// Inserting the request payload to the database
	if _, err := h.Col.InsertOne(context.TODO(), transcoder); err != nil {
		log.Errorf("Error while inserting the request: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	}

	// Returning the response
	return c.JSON(http.StatusCreated, "Transcoder added successfully.")
}
