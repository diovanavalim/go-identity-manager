package main

import (
	"fmt"
	"identity-manager/src/config"
	"identity-manager/src/router"
	"log"
	"net/http"
)

func main() {
	config.LoadEnvVariables()

	routes := router.CreateRouter()
	
	fmt.Println(fmt.Sprintf("Server Running on Port: %d", config.Port))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), routes))
}