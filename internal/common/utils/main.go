package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	OptionalAuth = false
	RequiredAuth = true
)

type M map[string]interface{}

func WriteJSON(w http.ResponseWriter, code int, data interface{}) {
	jsonBytes, err := json.Marshal(data)

	if err != nil {
		ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(jsonBytes)

	if err != nil {
		log.Println(err)
	}
}

func ReadJSON(body io.Reader, input interface{}) error {
	return json.NewDecoder(body).Decode(input)
}

func GenerateSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}
