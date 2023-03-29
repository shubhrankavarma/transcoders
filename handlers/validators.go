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
	if checkOutputInputType(i.(Transcoder).OutputType) == false {
		return errors.New("Invalid output type")
	}
	if checkOutputInputType(i.(Transcoder).InputType) == false {
		return errors.New("Invalid input type")
	}
	return tv.validator.Struct(i)
}
