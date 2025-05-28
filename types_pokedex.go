package main

type Pokemon struct {
	Name           string
	BaseExperience int
}

type Pokedex struct {
	caught map[string]Pokemon
}
