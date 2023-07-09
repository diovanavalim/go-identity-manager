package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}{
		Message:    err.Error(),
		StatusCode: statusCode,
	})
}
