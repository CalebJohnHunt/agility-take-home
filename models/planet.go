package models

import "time"

type Planet struct {
	Name           string    `json:"name"`
	RotationPeriod string    `json:"rotation_period"`
	OrbitalPeriod  string    `json:"orbital_period"`
	Diameter       string    `json:"diameter"`
	Climate        string    `json:"climate"`
	Gravity        string    `json:"gravity"`
	Terrain        string    `json:"terrain"`
	SurfaceWater   string    `json:"surface_water"`
	Population     string    `json:"population"`
	ResidentUrls   []string  `json:"residents"`
	FilmUrls       []string  `json:"films"`
	Created        time.Time `json:"created"`
	Edited         time.Time `json:"edited"`
	Url            string    `json:"url"`
}
