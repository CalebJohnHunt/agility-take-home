package models

import "time"

type Species struct {
	Name            string    `json:"name"`
	Classification  string    `json:"classification"`
	Designation     string    `json:"designation"`
	AverageHeight   string    `json:"average_height"`
	SkinColors      string    `json:"skin_colors"`
	HairColors      string    `json:"hair_colors"`
	EyeColors       string    `json:"eye_colors"`
	AverageLifespan string    `json:"average_lifespan"`
	HomeworldUrl    string    `json:"homeworld"`
	Language        string    `json:"language"`
	PeopleUrls      []string  `json:"people"`
	FilmUrls        []string  `json:"films"`
	Created         time.Time `json:"created"`
	Edited          time.Time `json:"edited"`
	Url             string    `json:"url"`
}
