package handlers

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func makeGoStructKey(key string) string {

	// Make character array from the key
	charArray := strings.Split(key, "")

	// Convert the first character to upper case
	charArray[0] = strings.ToUpper(charArray[0])

	// Make letter right after underscore to upper case
	for i := 1; i < len(charArray); i++ {
		if charArray[i] == "_" {
			charArray[i+1] = strings.ToUpper(charArray[i+1])
		}
	}

	newKey := strings.Join(charArray, "")

	// return cases.ToCamel(strings.ReplaceAll(newKey, "_", ""))
	return strings.Title(strings.ReplaceAll(newKey, "_", ""))
}

// PatchTranscoder is a handler to update a transcoder
// @description Update a transcoder
// @host localhost:51000
// @BasePath /
// @Router /transcoders [patch]
// @schemes http
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

	// OutputType and InputType from the query params
	outputType := c.QueryParam("output_type")
	inputType := c.QueryParam("input_type")

	// Check if the output type and input type is present in the query params
	if outputType == "" || inputType == "" {
		return c.JSON(http.StatusBadRequest, "Please provide output_type and input_type in query parameter.")
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
		if _, found := t.FieldByName(makeGoStructKey(key)); !found {
			return c.JSON(http.StatusBadRequest, "Invalid request payload.")
		}
	}

	// Filter for the update
	filter := bson.D{
		{Key: "output_type", Value: outputType},
		{Key: "input_type", Value: inputType},
	}

	// Update filter
	update := bson.D{}
	for key, val := range payload {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: key, Value: val}}})
	}

	// Options for the update
	opts := options.Update().SetUpsert(false)

	// Update the document
	if updated, err := h.Col.UpdateOne(c.Request().Context(), filter, update, opts); err != nil {
		log.Errorf("Error while updating the document: %v", err)
		return c.JSON(http.StatusInternalServerError, "Unable to update the document.")
	} else if updated.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, "No document found with the given output_type and input_type.")
	}

	// Return the response
	return c.JSON(http.StatusOK, "Document updated successfully.")
}
