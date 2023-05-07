package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/amagimedia/transcoders/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PatchTranscoder is a handler to update a transcoder
// @description Update a transcoder
// @Router /commands [patch]
// @Param input_type query string true "input_type" required=true
// @Param output_type query string true "output_type" required=true
// @Success 200 {object} string "Transcoder updated successfully."
// @Failure 400 {object} string "Please provide output_type and input_type in query parameter." example:"Please provide output_type and input_type in query parameter."
// @Failure 404 {object} string "Transcoder not found." example:"Transcoder not found."
// @Failure 500 {object} string "Unable to update the transcoder." example:"Unable to update the transcoder."
// @Failure 422 {object} string "Unable to pass the request payload." example:"Unable to pass the request payload."
// @Failure 400 {object} string "Invalid request payload." example:"Invalid request payload."
func (h *TranscoderHandler) PatchTranscoder(c echo.Context) error {

	// transcoder Variable
	var transcoder Transcoder

	// Filter to get the document
	filter, err := utils.MakeFilterUsingQueryParamToGetOneDocument(c)

	if err != nil {
		log.Error("Please provide asset_type and operation in query parameter.")
		return c.JSON(http.StatusBadRequest, "Please provide asset_type and operation in query parameter.")
	}

	// Reading the request payload in a map
	var payload map[string]interface{}
	if err := json.NewDecoder(c.Request().Body).Decode(&payload); err != nil {
		log.Errorf("Error while decoding the request: %v", err)
		return c.JSON(http.StatusUnprocessableEntity, "Unable to pass the request payload.")
	}

	// Check if the request payload has any invalid key
	t := reflect.TypeOf(transcoder)
	for key := range payload {
		if _, found := t.FieldByName(utils.MakeGoStructKey(key)); !found {
			log.Errorf("Invalid request payload.")
			return c.JSON(http.StatusBadRequest, "Invalid request payload.")
		}
	}

	// Update filter
	update := bson.D{}
	for key, val := range payload {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: key, Value: val}}})
	}

	// Options for the update - Not to create a new document if not found
	opts := options.Update().SetUpsert(false)

	// Update the document
	if updated, err := h.Col.UpdateOne(c.Request().Context(), filter, update, opts); err != nil {
		log.Errorf("Error while updating the document: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to update the document.")
	} else if updated.MatchedCount == 0 {
		log.Errorf("No document found with the given output_type and input_type.")
		return c.JSON(http.StatusNotFound, "No document found with the given output_type and input_type.")
	}

	// Return the response
	return c.JSON(http.StatusOK, "Document updated successfully.")
}
