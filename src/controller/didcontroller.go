package controller

import (
	"errors"
	"identity-manager/src/response"
	"net/http"
)

func CreateDID(w http.ResponseWriter, r *http.Request) {
	seed := r.URL.Query().Get("seed")

	if seed == "" || len(seed) != 32 {
		response.Error(w, http.StatusBadRequest, errors.New("invalid seed param"))
	}
}
