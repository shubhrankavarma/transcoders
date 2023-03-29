package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *TranscoderHandler) putTranscoder(c echo.Context) error {

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

	// Getting the ID from the request
	id := c.Param("id")

	// Getting the ID from the request
	filter := bson.D{{Key: "_id", Value: id}}

	// Update all filter
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "id", Value: id},
			{Key: "updated_at", Value: time.Now()},
			{Key: "output_type", Value: trancoder.OutputType},
			{Key: "input_type", Value: trancoder.InputType},
			{Key: "template_command", Value: trancoder.TemplateCommand},
			{Key: "updated_by", Value: trancoder.UpdatedBy},
		}},
	}

	// Options for the update
	opts := options.Update().SetUpsert(false)

	// Updating the request payload to the database
	if _, err := h.Col.UpdateOne(context.Background(), filter, update, opts); err != nil {
		log.Errorf("Error while updating the request: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	}

	// Returning the response
	return c.JSON(http.StatusOK, "Transcoder updated successfully.")
}
