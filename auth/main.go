package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"go-oauth2-server/auth/handler"
)

func main() {
	ah := handler.NewAuthorizeHandler()
	// th := handler.NewTokenHandler()

	http.HandleFunc("/authorize", ah.HandleAuthorizeRequest)
	// http.HandleFunc("/token", th.HandleTokenRequest)

	port := "9001"
	log.Printf("listen port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}
