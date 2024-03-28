package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/jamersom/golang-validator/utils"
	"time"
)

type ErrorResponse struct {
	Timestamp time.Time         `json:"timestamp"`
	Error     string            `json:"error"`
	Fields    []FieldValidation `json:"validations"`
}

type FieldValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ErrorResponse) AppendField(fe validator.FieldError) {
	message := utils.Error(fe)
	e.Fields = append(e.Fields, FieldValidation{
		Field:   fe.Field(),
		Message: message,
	})
}
