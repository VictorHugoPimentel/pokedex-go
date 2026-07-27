package service

import (
	"Pokedex/internal/pokeapi"
)

type LocationService struct {
	client pokeapi.Client
	nextURL *string
	prevURL *string
}

func NewLocationService(client pokeapi.Client) *LocationService {
	return &LocationService{
		client: client,
	}
}

func (s *LocationService) ExploreLocationArea(locationArea string) (pokeapi.LocationArea, error) {
	return s.client.GetLocationArea(locationArea)	
}
