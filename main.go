package main

import (
	"agility-take-home/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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
		panic(err)
	}

	people := searchResult.Results
	finals := make([]models.Final, len(people)) // Ahhh yeahh gotta save those allocations
	wg := sync.WaitGroup{}
	for i, p := range people {
		wg.Add(1)
		go func(i int, p models.Person) {
			defer wg.Done()
			wg1 := sync.WaitGroup{}
			finals[i].Person = p
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				starships, err := getFromUrls[models.Starship](p.StarshipUrls)
				if err != nil {
					panic(err)
				}
				finals[i].Starships = starships
			}()
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				homeplanet, err := getFromUrl[models.Planet](p.HomeworldUrl)
				if err != nil {
					panic(err)
				}
				finals[i].HomePlanet = homeplanet
			}()
			wg1.Add(1)
			go func() {
				defer wg1.Done()
				species, err := getFromUrls[models.Species](p.SpeciesUrls)
				if err != nil {
					panic(err)
				}
				finals[i].Species = species
			}()
			wg1.Wait()
		}(i, p)
	}
	wg.Wait()

	slices.SortFunc(finals, func(a, b models.Final) int { return strings.Compare(a.Person.Name, b.Person.Name) })

	for _, f := range finals {
		tmpl := template.Must(template.ParseGlob("templates/basic.tmpl"))
		if err = tmpl.Execute(os.Stdout, f); err != nil {
			panic(err)
		}
	}
}

func getFromUrls[T any](urls []string) ([]T, error) {
	ts := make([]T, len(urls))
	// errChan acts as our wait group
	errChan := make(chan error)
	defer close(errChan)
	for i, url := range urls {
		// Get each url concurrently
		go func(i int, url string) {
			t, err := getFromUrl[T](url)
			if err != nil {
				errChan <- err
				return
			}

			ts[i] = t
			errChan <- nil
		}(i, url)
	}
	var retErr error = nil
	for i := 0; i < len(urls); i++ {
		if err := <-errChan; err != nil {
			retErr = errors.Join(err, retErr)
		}
	}
	return ts, nil
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
