package models

// Character model of Rick and morty
type Character struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Gender string `json:"gender"`
}

// CharacterPage model for paginate API results
type CharacterPage struct {
	Info       `json:"info"`
	Characters []Character `json:"results"`
}
