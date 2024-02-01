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
	host := "0.0.0.0"
	log.Printf("listen start: %s:%s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
