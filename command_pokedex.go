package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex.caught) == 0 {
		fmt.Println("Your Pokedex is empty. Catch some Pokemon first!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for name := range c.pokedex.caught {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
