package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ErrorM is used to create the validation error response format according to the API spec
type ErrorM map[string][]string

// Error is needed to implement the error interface
func (e ErrorM) Error() string {
	return "validation error"
}

var (
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrDuplicateUsername = errors.New("duplicate username")
	ErrNotFound          = errors.New("record not found")
	ErrUnAuthorized      = errors.New("unauthorized")
	ErrInternal          = errors.New("internal error")
)

func ValidationError(w http.ResponseWriter, _err error) {
	resp := ErrorM{}

	switch err := _err.(type) {
	case validator.ValidationErrors:
		for _, e := range err {
			field := e.Field()
			msg := checkTagRules(e)
			resp[field] = append(resp[field], msg)
		}
	default:
		resp["non_field_error"] = append(resp["non_field_error"], err.Error())
	}
	ErrorResponse(w, http.StatusUnprocessableEntity, resp)
}

func BadRequestError(w http.ResponseWriter) {
	ErrorResponse(w, http.StatusUnprocessableEntity, "unable to process request")
}

func InvalidUserCredentialsError(w http.ResponseWriter) {
	msg := "invalid authentication credentials"
	ErrorResponse(w, http.StatusUnauthorized, msg)
}

func InvalidAuthTokenError(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Token")
	msg := "invalid or missing authentication token"
	ErrorResponse(w, http.StatusUnauthorized, msg)
}

func NotFoundError(w http.ResponseWriter, err ErrorM) {
	ErrorResponse(w, http.StatusNotFound, err)
}

func ServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	ErrorResponse(w, http.StatusInternalServerError, "internal error")
}

func ErrorResponse(w http.ResponseWriter, code int, errs interface{}) {
	WriteJSON(w, code, M{"errors": errs})
}

func checkTagRules(e validator.FieldError) (errMsg string) {
	tag, field, param, value := e.ActualTag(), e.Field(), e.Param(), e.Value()

	if tag == "required" {
		errMsg = "this field is required"
	}

	if tag == "email" {
		errMsg = fmt.Sprintf("%q is not a valid email", value)
	}

	if tag == "min" {
		errMsg = fmt.Sprintf("%s must be greater than %v", field, param)
	}

	if tag == "max" {
		errMsg = fmt.Sprintf("%s must be less than %v", field, param)
	}
	return
}
