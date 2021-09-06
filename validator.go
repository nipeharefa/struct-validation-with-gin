package main

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

const ValueIsInvalid string = "value is invalid"

var ExampleValidationTranslations map[string]string = map[string]string{
	"required":  ValueIsInvalid,
	"sep-comma": ValueIsInvalid,
}

func GetOneValidatorErrMsg(e validator.ValidationErrors) string {
	for _, v := range e {
		fmt.Println(v)
		return fmt.Sprintf("%s %s", v.Field(), ExampleValidationTranslations[v.Tag()])
	}

	return ""
}

func NewValidator() *validator.Validate {
	r := validator.New()

	r.RegisterValidation("is-awesome", ValidateMyVal)
	r.RegisterValidation("sep-comma", SepByComma)

	return r
}

func SepByComma(fl validator.FieldLevel) bool {
	gex := regexp.MustCompile("^([a-z]+)(,[a-z]+)*$")
	return gex.MatchString(fl.Field().String())
}

// ValidateMyVal implements validator.Func
func ValidateMyVal(fl validator.FieldLevel) bool {
	return fl.Field().String() == "awesome"
}
