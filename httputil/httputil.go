package httputil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
)

// HTTPError is sent with a non-200 code and error message
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// ParseQuery parses, decodes and validates query params
// It decodes the params into the obj struct using schema.Decoder
func ParseQuery(r *http.Request, obj any) error {
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("could not parse query: %w", err)
	}
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	if err := decoder.Decode(obj, r.Form); err != nil {
		return fmt.Errorf("could not decode query: %w", err)
	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return fmt.Errorf("missing or bad input: %w", err)
	}
	return nil
}

// ParseBody parses, decodes and validates a http.Request body
// It decodes the body into the obj struct using json.Unmarshal
func ParseBody(r *http.Request, obj any) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not read body: %w", err)
	}
	if err := json.Unmarshal(body, obj); err != nil {
		return fmt.Errorf("could not decode body: %w", err)
	}
	validate := validator.New()
	if err := validate.Struct(obj); err != nil {
		return fmt.Errorf("missing or bad input: %w", err)
	}
	return nil
}

// RespondWithJSON writes to a http.ResponseWriter with a json payload
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "could not marshal json response")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError writes to a http.ResponseWriter with an error string
func RespondWithError(w http.ResponseWriter, code int, message string) {
	if code == http.StatusInternalServerError {
		fmt.Println(message)
	}
	RespondWithJSON(w, code, HTTPError{Code: code, Message: message})
}
