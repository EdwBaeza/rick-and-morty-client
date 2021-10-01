package structs

// Character model of Rick and morty
type Character struct {
	Base
	Name    string   `json:"name"`
	Status  string   `json:"status"`
	Type    string   `json:"type"`
	Gender  string   `json:"gender"`
	Episode []string `json:"episode"`
}

// CharacterPage model for paginate API results
type CharacterPage struct {
	Info       `json:"info"`
	Characters []Character `json:"results"`
}
