package main

import "fmt"

func commandCatch(cfg *config, pokemon string) error {

	pokemonResp, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonResp.Name)
	fmt.Println(pokemonResp.BaseExperience)
	return nil
}
