package models

type Response[T any] struct {
	Count       int    `json:"count"`
	NextUrl     string `json:"next"`
	PreviousUrl string `json:"previous"`
	Results     []T    `json:"results"`
}
