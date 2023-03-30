package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var (
	v = validator.New()
)

type TranscoderValidator struct {
	validator *validator.Validate
}

func checkOutputInputType(tp string) bool {
	if _, ok := acceptedInputAndOutputTypes[tp]; ok {
		return true
	}
	return false
}

func (tv *TranscoderValidator) Validate(i interface{}) error {
	if !checkOutputInputType(i.(Transcoder).OutputType) {
		return errors.New("invalid output type")
	}
	if !checkOutputInputType(i.(Transcoder).InputType) {
		return errors.New("invalid input type")
	}
	if i.(Transcoder).InputType == i.(Transcoder).OutputType {
		return errors.New("input and output types cannot be same")
	}
	return tv.validator.Struct(i)
}
