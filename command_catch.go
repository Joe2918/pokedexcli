package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	res := rand.Intn(pokemonResp.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonResp.Name)
	fmt.Println(res)
	if res > 65 {
		fmt.Printf("%v escaped!\n", pokemonResp.Name)
		return nil
	}

	fmt.Printf("%v was caught!\n", pokemonResp.Name)

	cfg.caughtPokemon[pokemonResp.Name] = pokemonResp
	return nil
}
