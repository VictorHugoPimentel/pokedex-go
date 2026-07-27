package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	resp, err := cfg.locationService.NextLocations()
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println("-", area.Name)
	}
	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	resp, err := cfg.locationService.PreviousLocations()
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Println("-", area.Name)
	}
	return nil
}