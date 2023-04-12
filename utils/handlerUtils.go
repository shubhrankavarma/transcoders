package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeFilterToGetOneDocument(c echo.Context) (primitive.M, error) {

	// OutputType, InputType, Codec and descriptor from the query params
	outputType := c.QueryParam("output_type")
	inputType := c.QueryParam("input_type")
	codec := c.QueryParam("codec")
	descriptor := c.QueryParam("descriptor")

	// Check if the output type and input type is present in the query params
	if outputType == "" || inputType == "" || codec == "" || descriptor == "" {
		return nil, errors.New("please provide output_type, input_type, codec and descriptor in query parameter")
	}

	// Filter to delete the document
	filter := bson.M{
		"output_type": outputType,
		"input_type":  inputType,
		"codec":       codec,
		"descriptor":  descriptor,
	}

	return filter, nil
}
