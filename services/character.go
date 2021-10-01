package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/edwbaeza/rick-and-morty-client/structs"
	"github.com/go-resty/resty/v2"
)

type Character struct {
	http
}

// FindByID returns single character
func (loc *Character) FindByID(id int) (structs.Character, error) {
	character := structs.Character{}
	pathParams := map[string]string{"id": strconv.Itoa(id)}
	queryParams := map[string]string{}
	response, err := loc.get(loc.resource+"/{id}", pathParams, queryParams)

	if err != nil {
		log.Fatalln(fmt.Sprintf("FindByID %d, Error %s", id, err))
		return character, err
	}

	log.Println(fmt.Sprintf("Success Response FindByID Location: %d", id))
	json.Unmarshal(response, &character)
	return character, nil
}

// FindAll returns all characters by paging API
func (loc *Character) FindAll() ([]structs.CharacterPage, error) {
	var characterPages []structs.CharacterPage
	pathParams := map[string]string{}
	queryParams := map[string]string{}

	for i := 1; ; i++ {
		characterPage := structs.CharacterPage{}
		queryParams["page"] = strconv.Itoa(i)
		response, err := loc.get(loc.resource, pathParams, queryParams)

		if err != nil {
			log.Fatalln(fmt.Sprintf("FindAll at page %d, error: %s", i, err))
			return characterPages, err
		}

		json.Unmarshal(response, &characterPage)
		characterPages = append(characterPages, characterPage)

		if characterPage.Info.Next == "" {
			break
		}
	}

	return characterPages, nil
}

// Filter characters by following keys
// name: filter by the given name.
// status: filter by the given status (alive, dead or unknown).
// species: filter by the given species.
// type: filter by the given type.
// gender: filter by the given gender (female, male, genderless or unknown).
// func (c *Character) Filter(queryParams map[string]string) ([]structs.Character, error) {
func (loc *Character) Filter(params map[string]string) ([]structs.CharacterPage, error) {
	characterPages := []structs.CharacterPage{}
	pathParams := map[string]string{}

	for i := 1; ; i++ {
		characterPage := structs.CharacterPage{}
		params["page"] = strconv.Itoa(i)
		response, err := loc.get(loc.resource, pathParams, params)

		if err != nil {
			log.Fatalln(fmt.Sprintf("Filter at page %d, error: %s", i, err))
			return characterPages, err
		}

		json.Unmarshal(response, &characterPage)
		characterPages = append(characterPages, characterPage)

		log.Println(fmt.Sprintf("Success Response Filter Location Page: %d", i))

		if characterPage.Info.Next == "" {
			log.Println("Pagination finished")
			break
		}
	}

	return characterPages, nil
}

//NewCharacterService return new struct of character service
func NewCharacterService() *Character {
	return &Character{
		http: http{
			resource: "character",
			client:   resty.New(),
		},
	}
}
