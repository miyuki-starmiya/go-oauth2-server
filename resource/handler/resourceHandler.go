package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"go-oauth2-server/resource/service"
)

func ResourceHandler(w http.ResponseWriter, r *http.Request) {
	resource := service.GetResource(r)
	if resource == nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Unauthorized\n")
		log.Printf("Unauthorized\n")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resource)
	log.Printf("success\n")
}
