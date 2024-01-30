package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/authorize", authorizeHandler)
	http.HandleFunc("/token", tokenHandler)

	port := "9001"
	log.Printf("listen port: %s\n", port)
	http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}