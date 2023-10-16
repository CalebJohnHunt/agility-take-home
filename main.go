package main

import (
	"agility-take-home/models"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"text/template"
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
		log.Fatalf("Encountered error while searching for \"%s\". Error: %s", os.Args[1], err)
	}

	people := searchResult.Results
	finals := make([]models.Final, len(people)) // Ahhh yeahh gotta save those allocations
	wgPeople := sync.WaitGroup{}
	for i, p := range people {
		wgPeople.Add(1)
		go func(i int, p models.Person) {
			defer wgPeople.Done()
			fillFinalForPerson(p, &finals[i])
		}(i, p)
	}
	wgPeople.Wait()

	slices.SortFunc(finals, func(a, b models.Final) int { return strings.Compare(a.Person.Name, b.Person.Name) })

	outputFinals(finals)
}

// I'm not so sure about the out variable here, but I thought I'd give it a try.
func fillFinalForPerson(p models.Person, f *models.Final) {
	wgPerson := sync.WaitGroup{}
	f.Person = p
	wgPerson.Add(1)
	go func() {
		defer wgPerson.Done()
		starships, err := getFromUrls[models.Starship](p.StarshipUrls)
		if err != nil {
			panic(err)
		}
		f.Starships = starships
	}()
	wgPerson.Add(1)
	go func() {
		defer wgPerson.Done()
		homeplanet, err := getFromUrl[models.Planet](p.HomeworldUrl)
		if err != nil {
			panic(err)
		}
		f.HomePlanet = &homeplanet
	}()
	wgPerson.Add(1)
	go func() {
		defer wgPerson.Done()
		species, err := getFromUrls[models.Species](p.SpeciesUrls)
		if err != nil {
			panic(err)
		}
		f.Species = species
	}()
	wgPerson.Wait()
}

func outputFinals(finals []models.Final) {
	for _, f := range finals {
		tmpl := template.Must(template.ParseGlob("templates/basic.tmpl"))
		if err := tmpl.Execute(os.Stdout, f); err != nil {
			panic(err)
		}
	}
}
