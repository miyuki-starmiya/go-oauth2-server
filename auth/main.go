package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"go-oauth2-server/auth/handler"
)

func main() {
	http.HandleFunc("/authorize", handler.AuthorizeHandler)
	// http.HandleFunc("/token", handler.tokenHandler)

	port := "9001"
	log.Printf("listen port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}
