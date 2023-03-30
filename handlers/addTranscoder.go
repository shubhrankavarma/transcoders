package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

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
