package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
)

func (h *TranscoderHandler) DeleteTranscoder(c echo.Context) error {

	// OutputType and InputType from the query params
	outputType := c.QueryParam("output_type")
	inputType := c.QueryParam("input_type")

	// Check if the output type and input type is present in the query params
	if outputType == "" || inputType == "" {
		log.Error("Please provide output_type and input_type in query parameter.")
		return c.JSON(http.StatusBadRequest, "Please provide output_type and input_type in query parameter.")
	}

	// Filter to delete the document
	filter := bson.M{
		"output_type": outputType,
		"input_type":  inputType,
	}

	// Deleting the document from the database
	if deleted, err := h.Col.DeleteOne(c.Request().Context(), filter); err != nil {
		log.Error("Unable to delete the transcoder.", err)
		return c.JSON(http.StatusInternalServerError, "Unable to delete the transcoder.")
	} else if deleted.DeletedCount == 0 {
		log.Error("Transcoder not found.")
		return c.JSON(http.StatusNotFound, "Transcoder not found.")
	}

	// Return the response
	return c.JSON(http.StatusOK, "Transcoder deleted successfully.")
}
