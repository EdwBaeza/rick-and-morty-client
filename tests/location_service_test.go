package tests

import (
	"testing"

	"github.com/edwbaeza/rick-and-morty-client/services"
)

func TestLocationFindbyId(t *testing.T) {
	const id = 1
	locationService := services.NewLocationService()
	location, err := locationService.FindByID(1)

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if location.ID != id {
		t.Error("Invalid found location")
	}
}

func TestLocationFindAll(t *testing.T) {
	locationService := services.NewLocationService()
	locationPages, err := locationService.FindAll()

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if len(locationPages) != 6 {
		t.Errorf("Invalid locations pages: %d", len(locationPages))
	}
}

func TestLocationFilter(t *testing.T) {
	queryParams := map[string]string{
		"name":   "rick",
		"status": "alive",
	}
	locationService := services.NewLocationService()
	filteredLocationPages, err := locationService.Filter(queryParams)

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if len(filteredLocationPages) != 1 {
		t.Errorf("Invalid locations pages: %d", len(filteredLocationPages))
	}
}
