package models

import "time"

type Person struct {
	Name         string    `json:"name"`
	Height       string    `json:"height"`
	Mass         string    `json:"mass"`
	HairColor    string    `json:"hair_color"`
	SkinColor    string    `json:"skin_color"`
	EyeColor     string    `json:"eye_color"`
	BirthYear    string    `json:"birth_year"`
	Gender       string    `json:"gender"`
	HomeworldUrl string    `json:"homeworld"`
	SpeciesUrls  []string  `json:"species"`
	VehiclesUrls []string  `json:"vehicles"`
	StarshipUrls []string  `json:"starships"`
	FilmUrls     []string  `json:"films"`
	Created      time.Time `json:"created"`
	Edited       time.Time `json:"edited"`
	Url          string    `json:"url"`
}
