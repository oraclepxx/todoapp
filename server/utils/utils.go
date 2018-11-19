package utils

import (
	"encoding/json"
	"net/http"
	"errors"

	uuid "github.com/satori/go.uuid"
)

func WriteJSON(w http.ResponseWriter, code int, value interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(value)
}

func GenerateUUID() (string, error) {
	uuidv4, err := uuid.NewV4()
	if err != nil {
		return "", errors.New(err.Error())
	}
	return uuidv4.String(), nil
}
