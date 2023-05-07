package utils

import (
	"errors"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var requiredParams []string = []string{"asset_type", "operation"}

func getParamAndGenerateFilter(c echo.Context, filter *bson.M, params []string) error {

	for _, param := range params {
		paramValue := c.QueryParam(param)
		if paramValue == "" {
			return errors.New("please provide " + param + " in query parameter")
		}
		(*filter)[param] = paramValue
	}

	return nil
}

func MakeFilterUsingQueryParamToGetOneDocument(c echo.Context) (primitive.M, error) {

	// Filter
	filter := bson.M{}

	// Get the required query params
	err := getParamAndGenerateFilter(c, &filter, requiredParams)
	if err != nil {
		return nil, err
	}

	// If asset_type is audio and operation is extraction, then check for these params -
	// audioCount, channelsOneCount, channelsTwoCount, channelsSixCount, channelsEightCount
	if filter["asset_type"] == "audio" && filter["operation"] == "extraction" {

		audioExtractionParams := []string{"audio_count", "channels_one_count", "channels_two_count", "channels_six_count", "channels_eight_count"}

		// Update the filter
		err := getParamAndGenerateFilter(c, &filter, audioExtractionParams)
		if err != nil {
			return nil, err
		}

	}

	// If asset_type is video and operation is extraction, then check for these params -
	// inputScanType, outputScanType
	if filter["asset_type"] == "video" && filter["operation"] == "extraction" {

		videoExtractionParams := []string{"input_scan_type", "output_scan_type"}

		// Update the filter
		err := getParamAndGenerateFilter(c, &filter, videoExtractionParams)
		if err != nil {
			return nil, err
		}
	}

	return filter, nil
}
