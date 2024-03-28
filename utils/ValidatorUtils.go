package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	FIELD_MSG_DEFAULT = "Validation: Field '%s' failed on the '%s = %s'"

	FIELD_MSG_MIN = "Validation: Field '%s' failed on the %s = '%s'"
	FIELD_MSG_MAX = "Validation: Field '%s' failed on the '%s = %s'"

	FIELD_MSG_GREATE_THEN = "Validation: Field '%s' failed greater than or equal = '%s'"
	FIELD_MSG_LESS_THEN   = "Validation: Field '%s' failed less than or equal = '%s'"
	FIELD_MSG_CHOOSE      = "Validation: Field '%s' failed when choosing between '%s'"
	FIELD_MSG_EMAIL       = "Validation: Field '%s' is not invalid"
)

func Error(fe validator.FieldError) string {
	if fe.Tag() == "email" {
		return fmt.Sprintf(FIELD_MSG_EMAIL, fe.Field())
	}
	if fe.Tag() == "oneof" {
		return fmt.Sprintf(FIELD_MSG_CHOOSE, fe.Field(), fe.Param())
	}
	if fe.Tag() == "gte" {
		return fmt.Sprintf(FIELD_MSG_GREATE_THEN, fe.Field(), fe.Param())
	}
	if fe.Tag() == "lte" {
		return fmt.Sprintf(FIELD_MSG_LESS_THEN, fe.Field(), fe.Param())
	}
	if fe.Tag() == "min" {
		return fmt.Sprintf(FIELD_MSG_MIN, fe.Field(), fe.Tag(), fe.Param())
	}
	if fe.Tag() == "max" {
		return fmt.Sprintf(FIELD_MSG_MAX, fe.Field(), fe.Tag(), fe.Param())
	}
	if fe.Tag() == "teste" {
		fmt.Sprintf("TAG :", fe.Tag())
		return fmt.Sprintf(FIELD_MSG_MAX, fe.Field(), fe.Tag(), fe.Param())
	}

	return fmt.Sprintf(FIELD_MSG_DEFAULT, fe.Field(), fe.Tag(), fe.Param())
}
