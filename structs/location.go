package structs

// Location model of Rick and morty
type Location struct {
	Base
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Dimension string   `json:"dimension"`
	URL       string   `json:"url"`
	Residents []string `json:"residents"`
}

// LocationPage model for paginate API results
type LocationPage struct {
	Info      `json:"info"`
	Locations []Location `json:"results"`
}
