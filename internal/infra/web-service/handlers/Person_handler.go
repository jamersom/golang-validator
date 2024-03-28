package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jamersom/golang-validator/internal/request"
	"github.com/jamersom/golang-validator/internal/response"
	"log/slog"
	"net/http"
	"time"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

type PersonHandler struct {
}

func NewPersonHandler() *PersonHandler {
	return &PersonHandler{}
}

func (h PersonHandler) PersonCreateHandler(w http.ResponseWriter, r *http.Request) {
	request := &request.PersonRequst{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		slog.Error("error ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	validationError := validates(request)
	if validationError != nil {

		errResp := response.ErrorResponse{
			Timestamp: time.Now(),
			Error:     "StatusBadRequest",
		}
		for _, fieldError := range validationError {
			errResp.AppendField(fieldError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errResp)

	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func validates(request *request.PersonRequst) validator.ValidationErrors {
	validate = validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(request)

	if err != nil {
		validationErrors, ok := err.(validator.ValidationErrors)
		if ok {
			for _, validationError := range validationErrors {
				fmt.Println(validationError.Field(), validationError.Tag(), validationError.Param())
			}
			return validationErrors
		} else {
			slog.Error("error validator Person:", err)
		}
		return nil
	}
	return nil
}
