package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokemon name")
	}

	name := args[0]
	pokemon, err := c.pokeapiClient.CatchPokemon(name)
	if err != nil {
		return err
	}

	roll := rand.Float64()
	chance := (315.0 - float64(pokemon.BaseExperience)) / 260.0
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if roll < chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		c.pokedex.caught[pokemon.Name] = Pokemon{pokemon.Name, pokemon.BaseExperience}
		fmt.Printf("You may now inspect it with the inspect command.\n")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
