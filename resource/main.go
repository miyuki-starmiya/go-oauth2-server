package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"go-oauth2-server/resource/handler"
)

func main() {
	rh := handler.NewResourceHandler()
	http.HandleFunc("/resource", rh.GetResource)

	port := "9002"
	log.Printf("listen port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}
