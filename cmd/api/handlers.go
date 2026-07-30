package main

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message"`
}

func (s *Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	response := PingResponse{
		Message: "pong",
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}