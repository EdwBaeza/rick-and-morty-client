package tests

import (
	"testing"

	"github.com/edwbaeza/rick-and-morty-client/services"
)

func TestCharacterFindbyId(t *testing.T) {
	const id = 1
	characterService := services.NewCharacterService()
	character, err := characterService.FindByID(1)

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if character.ID != id {
		t.Error("Invalid found character")
	}
}

func TestCharacterFindAll(t *testing.T) {
	characterService := services.NewCharacterService()
	characterPages, err := characterService.FindAll()

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if len(characterPages) != 34 {
		t.Error("Invalid characters pages")
	}
}

func TestCharacterFilter(t *testing.T) {
	queryParams := map[string]string{
		"name":   "rick",
		"status": "alive",
	}
	characterService := services.NewCharacterService()
	filteredCharacterPages, err := characterService.Filter(queryParams)

	if err != nil {
		t.Errorf(string(err.Error()))
	}

	if len(filteredCharacterPages) != 2 {
		t.Error("Invalid characters pages")
	}
}
