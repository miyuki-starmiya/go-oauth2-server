package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	"github.com/miyuki-starmiya/go-oauth2-server/db/store"
	"github.com/miyuki-starmiya/go-oauth2-server/resource/handler"
)

func main() {
	db, err := store.NewDatabase()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}

	rh := handler.NewResourceHandler(
		store.NewTokenStore(db),
	)
	http.HandleFunc("/resource", rh.HandleResourceRequest)

	port := "9002"
	host := "0.0.0.0"
	log.Printf("listen start: %s:%s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
