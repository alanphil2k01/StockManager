package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func ParseBody(r *http.Request, x interface{}) {
	if err := json.NewDecoder(r.Body).Decode(x); err != nil {
		return
	}
}
