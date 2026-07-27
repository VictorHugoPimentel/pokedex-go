package service

import (
	"Pokedex/internal/pokeapi"
	"errors"
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

func (s *LocationService) NextLocations() (pokeapi.LocationAreasResp, error) {
	resp, err := s.client.ListLocationAreas(s.nextURL)
	if err != nil {
		return pokeapi.LocationAreasResp{}, err
	}

	s.nextURL = resp.Next
	s.prevURL = resp.Previous

	return resp, nil
}

func (s *LocationService) PreviousLocations() (pokeapi.LocationAreasResp, error) {
	if s.prevURL == nil {
		return pokeapi.LocationAreasResp{}, errors.New("you are on the first page")
	}

	resp, err := s.client.ListLocationAreas(s.prevURL)
	if err != nil {
		return pokeapi.LocationAreasResp{}, err
	}

	s.nextURL = resp.Next
	s.prevURL = resp.Previous

	return resp, nil
}