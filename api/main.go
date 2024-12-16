package main

import (
	"log"
	"net/http"
	"server/routers"
)

func main() {
	router := routers.SetupRouter()
	log.Println("APP started and running on PORT 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
