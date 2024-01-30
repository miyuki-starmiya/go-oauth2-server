package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go-oauth2-server/resource/service"
)

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	resource, err := service.GetResource(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Unauthorized\n")
		log.Printf("error: %v\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
	log.Printf("success\n")
}
