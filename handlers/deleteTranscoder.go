package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteTranscoder is a handler to delete a transcoder
// @description Delete a transcoder
// @host localhost:51000
// @BasePath /
// @Router /transcoders [delete]
// @schemes http
// @Param Authorization header string true "JWT Token"
// @Param input_type query string true "input_type" required=true
// @Param output_type query string true "output_type" required=true
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string "Transcoder Delted successfully."
// @Failure 400 {object} string "Please provide output_type and input_type in query parameter." example:"Please provide output_type and input_type in query parameter."
// @Failure 404 {object} string "Transcoder not found." example:"Transcoder not found."
// @Failure 500 {object} string "Unable to delete the transcoder." example:"Unable to delete the transcoder."
func (h *TranscoderHandler) DeleteTranscoder(c echo.Context) error {

	// OutputType, InputType and Codec from the query params
	outputType := c.QueryParam("output_type")
	inputType := c.QueryParam("input_type")
	codec := c.QueryParam("codec")
	descriptor := c.QueryParam("descriptor")

	// Check if the output type and input type is present in the query params
	if outputType == "" || inputType == "" || codec == "" || descriptor == "" {
		log.Error("Please provide output_type, input_type, codec and descriptor in query parameter.")
		return c.JSON(http.StatusBadRequest, "Please provide output_type and input_type in query parameter.")
	}

	// Filter to delete the document
	filter := bson.M{
		"output_type": outputType,
		"input_type":  inputType,
		"codec":       codec,
		"descriptor":  descriptor,
		"status":      Active,
	}

	// Update - Set the status to inactive
	update := bson.M{
		"$set": bson.M{
			"status": "inactive",
		},
	}

	// Deleting the document from the database (updating the status to inactive)
	if deleted, err := h.Col.UpdateOne(c.Request().Context(), filter, update); err != nil {
		log.Error("Unable to delete the transcoder.", err)
		return c.JSON(http.StatusInternalServerError, "Unable to delete the transcoder.")
	} else if deleted.MatchedCount == 0 {
		log.Error("Transcoder not found.")
		return c.JSON(http.StatusNotFound, "Transcoder not found.")
	}

	// Return the response
	return c.JSON(http.StatusOK, "Transcoder deleted successfully.")
}
