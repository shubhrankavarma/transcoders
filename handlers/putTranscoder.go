package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PutTranscoder is a handler function to update the transcoder
// @Summary Update the transcoder
// @Description Update the transcoder
// @host localhost:51000
// @BasePath /
// @Router /transcoders [put]
// @schemes http
// @Param id query string true "id" required=true
// @Success 200 {object} string "Transcoder updated successfully."
// @Failure 400 {object} string "Please provide id in query parameter." example:"Please provide id in query parameter."
// @Failure 404 {object} string "Transcoder not found." example:"Transcoder not found."
// @Failure 500 {object} string "Unable to update the transcoder." example:"Unable to update the transcoder."
// @Failure 422 {object} string "Unable to pass the request payload." example:"Unable to pass the request payload."
func (h *TranscoderHandler) PutTranscoder(c echo.Context) error {

	// For Filter Paramters - Codec, Descriptor

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

	// Creating filter with output_type and input_type
	filter := bson.M{
		"output_type": transcoder.OutputType,
		"input_type":  transcoder.InputType,
		"codec":       transcoder.Codec,
		"descriptor":  transcoder.Descriptor,
	}

	// Options for the update - Not to create a new document if not found
	opts := options.Update().SetUpsert(false)

	// Marshalling the request payload
	// transcoderBSON, err := bson.Marshal(transcoder)
	// if err != nil {
	// 	log.Errorf("Error while marshalling the request: %v", err)
	// 	return c.JSON(http.StatusInternalServerError, "Unable to process the request.")
	// }

	// Update filter
	update := bson.D{
		{Key: "$set", Value: transcoder},
	}

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
