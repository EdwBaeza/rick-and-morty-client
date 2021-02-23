package models

import "time"

// Location model of Rick and morty
type Location struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	URL       string    `json:"url"`
	Residents []string  `json:"residents"`
	Created   time.Time `json:"created"`
}

// LocationPage model for paginate API results
type LocationPage struct {
	Info      `json:"info"`
	Locations []Location `json:"results"`
}
