package main

import (
	"app/internal/service"
	"fmt"
	"log"

	"net/http"
)

func main() {
	port := "8080"
	log.Printf("Server on %s is starting", port)
	service.Handlers()
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

}
