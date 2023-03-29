package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *TranscoderHandler) PutTranscoder(c echo.Context) error {

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
	id := c.QueryParam("id")

	// Getting the ID from the request
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Error while converting the id to object id: %v", err)
		return c.JSON(http.StatusBadRequest, "Invalid id.")
	}
	filter := bson.D{{Key: "_id", Value: objectId}}

	// Update all filter
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "updated_at", Value: time.Now()},
			{Key: "output_type", Value: trancoder.OutputType},
			{Key: "input_type", Value: trancoder.InputType},
			{Key: "template_command", Value: trancoder.TemplateCommand},
			{Key: "updated_by", Value: trancoder.UpdatedBy},
			{Key: "status", Value: trancoder.Status},
		}},
	}

	// Options for the update
	opts := options.Update().SetUpsert(false)

	// Updating the request payload to the database
	if r, err := h.Col.UpdateOne(context.Background(), filter, update, opts); err != nil {
		log.Errorf("Error while updating the request: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	} else if r.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, "Transcoder not found with the given id.")
	}

	// Returning the response
	return c.JSON(http.StatusOK, "Transcoder updated successfully.")
}
