package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	location, err := c.pokeapiClient.ExploreLocations(name)
	if err != nil {
		return err
	}

	fmt.Println("Exploring", location)
	fmt.Println("Found Pokemon:")
	for _, encounter := range location.PokemonEncounters {
		fmt.Println(" - ", encounter.Pokemon.Name)
	}
	return nil
}
