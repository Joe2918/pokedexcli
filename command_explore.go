package main

import "fmt"

func commandExplore(cfg *config, args string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(args)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + args)
	fmt.Println("Found Pokemon:")
	for _, poke := range pokemonResp.PokemonEncounters {
		fmt.Println(" - " + poke.Pokemon.Name)
	}
	return nil
}
