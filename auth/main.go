package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"go-oauth2-server/auth/handler"
	"go-oauth2-server/auth/store"
)

func main() {
	// store := store.NewStore()
	db, err := store.NewDatabase()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	ah := handler.NewAuthorizeHandler(
		store.NewCodeStore(db),
	)
	th := handler.NewTokenHandler(
		store.NewCodeStore(db),
		// add TokenStore
	)

	http.HandleFunc("/authorize", ah.HandleAuthorizeRequest)
	http.HandleFunc("/token", th.HandleTokenRequest)

	port := "9001"
	host := "0.0.0.0"
	log.Printf("listen start: %s:%s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
