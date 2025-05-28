package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, args ...string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(c *config, args ...string) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := c.pokeapiClient.ListLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = locationsResp.Next
	c.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
