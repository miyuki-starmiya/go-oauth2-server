package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"go-oauth2-server/resource/handler"
)

func main() {
	http.HandleFunc("/resource", handler.ResourceHandler)

	port := "9002"
	log.Printf("listen port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}
