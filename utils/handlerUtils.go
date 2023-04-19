package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeFilterUsingQueryParamToGetOneDocument(c echo.Context) (primitive.M, error) {

	assetType := c.QueryParam("asset_type")
	operation := c.QueryParam("operation")

	// Check if the output type and input type is present in the query params
	if operation == "" || assetType == "" {
		return nil, errors.New("please provide asset_type and operation in query parameter")
	}

	// Filter to delete the document
	filter := bson.M{
		"asset_type": assetType,
		"operation":  operation,
	}

	return filter, nil
}
