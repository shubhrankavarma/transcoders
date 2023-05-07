package handlers

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var (
	v = validator.New()
)

type TranscoderValidator struct {
	validator *validator.Validate
}

func (tv *TranscoderValidator) Validate(i interface{}) error {

	transcoder := i.(Transcoder)

	// If asset_type is audio and operation is extraction, then check for these params -
	// audioCount, channelsOneCount, channelsTwoCount, channelsSixCount, channelsEightCount
	if transcoder.AssetType == "audio" && transcoder.Operation == "extraction" {

		audioExtractionFields := []string{"AudioCount", "ChannelsOneCount", "ChannelsTwoCount", "ChannelsSixCount", "ChannelsEightCount"}

		// Check for these params
		for _, field := range audioExtractionFields {

			// Check for required tag using reflect
			field, found := reflect.TypeOf(transcoder).FieldByName(field)

			if !found {
				return errors.New(field.Name + " not found")
			}

			// If field is null, then return error
			if reflect.ValueOf(transcoder).FieldByName(field.Name).IsNil() {
				return errors.New(field.Name + " is required")
			}
		}

	}

	// If asset_type is video and operation is extraction, then check for these params -
	// inputScanType, outputScanType
	if transcoder.AssetType == "video" && transcoder.Operation == "extraction" {

		videoExtractionFields := []string{"InputScanType", "OutputScanType"}

		// Check for these params
		for _, field := range videoExtractionFields {

			// Check for required tag using reflect
			field, found := reflect.TypeOf(transcoder).FieldByName(field)

			if !found {
				return errors.New(field.Name + " not found")
			}

			// If field is null, then return error
			if reflect.ValueOf(transcoder).FieldByName(field.Name).IsNil() {
				return errors.New(field.Name + " is required")
			}
		}

	}

	return tv.validator.Struct(i)
}
