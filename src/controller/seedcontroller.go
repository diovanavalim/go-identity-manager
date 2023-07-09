package controller

import (
	"fmt"
	"identity-manager/src/dto"
	"identity-manager/src/response"
	"identity-manager/src/service/anchorhandle"
	"net/http"
)

func CreateSeed(w http.ResponseWriter, r *http.Request) {
	service := anchorhandle.Service{}

	seed := service.GenerateSeed()
	key := service.GenerateKey(48)

	fmt.Println(fmt.Sprintf("seed %s", seed))
	fmt.Println(fmt.Sprintf("key %s", key))

	response.JSON(w, http.StatusOK, dto.Seed{
		Seed: seed,
		Key:  key,
	})
}
