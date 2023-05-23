/*
Copyright © 2023 Xavier Portilla Edo

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

const pokemons = 151
const datasetFile = "dataset.jsonl"

type PokemonVertexDatasetLine struct {
	InputText  string `json:"input_text"`
	OutputText string `json:"output_text"`
}

func main() {

	// Create the file
	f, err := createFile()
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		return
	}
	defer f.Close()

	// Iterate over pokemons
	for id := 1; id <= pokemons; id++ {
		pokemon, err := pokeapi.Pokemon(strconv.Itoa(id))
		if err != nil {
			log.Fatalf("Error: %v\n", err)
			return
		}

		// create the Pokemon information
		pokemonJsonLine, err := createPokemonLine(pokemon)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
			return
		}

		// append the pokemon to the final dataset
		log.Printf("Writing: %s\n", pokemon.Name)
		if _, err := f.WriteString(pokemonJsonLine + "\n"); err != nil {
			log.Fatalf("Error: %v\n", err)
			return
		}
	}
	log.Printf("Process finished without errors\n")
}

func createPokemonLine(pokemon structs.Pokemon) (string, error) {

	species, err := pokeapi.PokemonSpecies(pokemon.Species.Name)
	if err != nil {
		return "", err
	}

	pokemonDescription := ""
	for _, spec := range species.FlavorTextEntries {
		if strings.Contains(spec.Language.Name, "en") {
			pokemonDescription = strings.ReplaceAll(spec.FlavorText, "\n", "")
		}
	}

	pokemonJsonLine := &PokemonVertexDatasetLine{InputText: pokemon.Name, OutputText: pokemonDescription}

	pokemonMarshalled, err := json.Marshal(pokemonJsonLine)
	if err != nil {
		return "", err
	}

	return string(pokemonMarshalled), nil
}

func createFile() (*os.File, error) {
	// Remove the file if exists
	if _, err := os.Stat(datasetFile); !errors.Is(err, os.ErrNotExist) {
		if err := os.Remove(datasetFile); err != nil {
			log.Fatalf("Error: %v\n", err)
			return nil, err
		}
	}
	log.Printf("Creating file: %v\n", datasetFile)

	// Open the file
	f, err := os.OpenFile(datasetFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error: %v\n", err)
		return nil, err
	}

	return f, nil
}
