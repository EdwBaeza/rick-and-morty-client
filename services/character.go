package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/edwbaeza/rick-and-morty-client/models"
	"github.com/go-resty/resty/v2"
)

type Character struct {
	http
}

// FindByID returns single character
func (c *Character) FindByID(id int) (models.Character, error) {
	character := models.Character{}
	response, err := c.http.client.R().EnableTrace().Get(c.http.FullPathWithID(id))

	if err != nil {
		log.Fatalln(fmt.Sprintf("FindByID %d, Error %s", id, err))
		return character, err
	}

	log.Println(fmt.Sprintf("Success Response FindByID Character : %d", id))
	json.Unmarshal(response.Body(), &character)
	return character, nil
}

// FindAll returns all characters by paging API
func (c *Character) FindAll() ([]models.CharacterPage, error) {
	characterPages := []models.CharacterPage{}
	url := c.http.FullPath()

	for i := 1; ; i++ {
		characterPage := models.CharacterPage{}
		queryParams := map[string]string{"page": strconv.Itoa(i)}

		response, err := c.http.client.R().
			SetQueryString(c.http.GenerateQueryParams(queryParams)).
			EnableTrace().
			Get(url)

		if err != nil {
			log.Fatalln(fmt.Sprintf("FindAll at page %d, error: %s", i, err))
			return characterPages, err
		}

		json.Unmarshal(response.Body(), &characterPage)
		characterPages = append(characterPages, characterPage)

		log.Println(fmt.Sprintf("Success Response FindAll Character Page: %d", i))

		if characterPage.Info.Next == "" {
			log.Println("Pagination finished")
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
func (c *Character) Filter(queryParams map[string]string) ([]models.CharacterPage, error) {
	characterPages := []models.CharacterPage{}
	url := c.http.FullPath()

	for i := 1; ; i++ {
		characterPage := models.CharacterPage{}
		queryParams["page"] = strconv.Itoa(i)

		response, err := c.http.client.R().
			SetQueryString(c.http.GenerateQueryParams(queryParams)).
			EnableTrace().
			Get(url)

		if err != nil {
			log.Fatalln(fmt.Sprintf("Filter at page %d, error: %s", i, err))
			return characterPages, err
		}

		json.Unmarshal(response.Body(), &characterPage)
		characterPages = append(characterPages, characterPage)

		log.Println(fmt.Sprintf("Success Response Filter Character Page: %d", i))

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
			relativePath: "character",
			client:       resty.New(),
		},
	}
}
