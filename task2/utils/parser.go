package utils

import (
	"net/http"
	"strconv"
)

func QueryParser(w http.ResponseWriter, r *http.Request, attr string) float64 {
	parsedValue, err := strconv.ParseFloat(r.URL.Query().Get(attr), 64)
	if err != nil {
		return -1
	}
	return parsedValue
}
