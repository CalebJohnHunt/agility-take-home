package models

import "time"

type Starship struct {
	Name                 string    `json:"name"`
	Model                string    `json:"model"`
	CostInCredits        string    `json:"cost_in_credits"`
	Length               string    `json:"length"`
	MaxAtmospheringSpeed string    `json:"max_atmosphering_speed"`
	Crew                 string    `json:"crew"`
	PassengerCount       string    `json:"passengers"`
	CargoCapacity        string    `json:"cargo_capacity"`
	Consumables          string    `json:"consumables"`
	HyperdriveRating     string    `json:"hyperdrive_rating"`
	MGLT                 string    `json:"MGLT"`
	StarshipClass        string    `json:"starship_class"`
	PilotUrls            []string  `json:"pilots"`
	FilmUrls             []string  `json:"films"`
	Created              time.Time `json:"created"`
	Edited               time.Time `json:"edited"`
	Url                  string    `json:"url"`
}
