package models

type Final struct {
	Person     Person
	Starships  []Starship
	// This is a pointer to allow for the possibility of no HomePlanet
	HomePlanet *Planet
	Species    []Species
}
