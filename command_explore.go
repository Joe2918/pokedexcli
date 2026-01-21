package main

import "fmt"

func commandExplore(cfg *config, args string) error {
	locationResp, err := cfg.pokeapiClient.ListPokemon(args)
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
