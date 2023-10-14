package main

import (
	"agility-take-home/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"

	"slices"
)

const (
	baseUri = "https://swapi.dev"
	apiUri  = baseUri + "/api"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("./%s <search>", os.Args[0])
		return
	}

	searchResult, err := getFromUrl[models.Response[models.Person]](fmt.Sprintf("%s/people?search=%s", apiUri, os.Args[1]))
	if err != nil {
		panic(err)
	}

	people := searchResult.Results
	finals := make([]models.Final, len(people)) // Ahhh yeahh gotta save those allocations
	for i, p := range people {
		finals[i].Person = p
		starships, err := getFromUrls[models.Starship](p.StarshipUrls)
		if err != nil {
			panic(err)
		}
		finals[i].Starships = starships
		homeplanet, err := getFromUrl[models.Planet](p.HomeworldUrl)
		if err != nil {
			panic(err)
		}
		finals[i].HomePlanet = homeplanet
		species, err := getFromUrls[models.Species](p.SpeciesUrls)
		if err != nil {
			panic(err)
		}
		finals[i].Species = species
	}

	slices.SortFunc(finals, func(a, b models.Final) int { return strings.Compare(a.Person.Name, b.Person.Name) })

	for _, f := range finals {
		tmpl := template.Must(template.ParseGlob("templates/basic.tmpl"))
		if err = tmpl.Execute(os.Stdout, f); err != nil {
			panic(err)
		}
	}
}

func getFromUrls[T any](urls []string) ([]T, error) {
	starships := make([]T, len(urls))
	for i, url := range urls {
		starship, err := getFromUrl[T](url)
		if err != nil {
			return nil, err
		}

		starships[i] = starship
	}
	return starships, nil
}

func getFromUrl[T any](url string) (T, error) {
	var t T = *new(T)
	resp, err := http.Get(url)
	if err != nil {
		return t, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		return t, err
	}

	err = json.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}
