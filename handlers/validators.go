package handlers

import "github.com/go-playground/validator/v10"

var (
	v = validator.New()
)

type TranscoderValidator struct {
	validator *validator.Validate
}

func (tv *TranscoderValidator) Validate(i interface{}) error {
	return tv.validator.Struct(i)
}
