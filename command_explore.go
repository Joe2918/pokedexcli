package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	name := args[0]
	locationResp, err := cfg.pokeapiClient.ListPokemon(name)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + locationResp.Location.Name)
	fmt.Println("Found Pokemon:")
	for _, poke := range locationResp.PokemonEncounters {
		fmt.Println(" - " + poke.Pokemon.Name)
	}
	return nil
}
