package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"go-oauth2-server/resource-server/usecase"
)

func main() {
    http.HandleFunc("/resource", resourceHandler)
    port := "9002"
    log.Printf("listen port: %s\n", port)
    http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil)
}

func resourceHandler(w http.ResponseWriter, r *http.Request) {
    resource := usecase.GetResource(r)
    if resource == nil {
        w.WriteHeader(http.StatusUnauthorized)
        io.WriteString(w, "Unauthorized\n")
        log.Printf("Unauthorized\n")
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resource)
    log.Printf("success\n")
}
