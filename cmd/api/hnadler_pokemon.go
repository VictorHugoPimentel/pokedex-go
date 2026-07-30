package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (s *Server) pokemonHandler(
	w http.ResponseWriter,
	r *http.Request,
) {

	name := strings.TrimPrefix(
		r.URL.Path,
		"/pokemon/",
	)

	pokemon, err := s.pokemonService.GetPokemon(name)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(pokemon)
}