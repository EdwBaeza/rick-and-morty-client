package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/edwbaeza/rick-and-morty-client/models"
	"github.com/go-resty/resty/v2"
)

//Character wrap for api
type Character struct {
	http
}

// FindByID return single character
func (c *Character) FindByID(id int) (models.Character, error) {
	character := models.Character{}
	response, err := c.http.client.R().EnableTrace().Get(c.http.FullPathWithID(id))

	if err != nil {
		log.Fatalln(fmt.Sprintf("FindByID %d, Error %s", id, err))
		return character, err
	}

	log.Println(fmt.Sprintf("Success Response Character FindByID: %d", id))
	json.Unmarshal(response.Body(), &character)
	return character, nil
}

// FindAll return all character
func (c *Character) FindAll() ([]models.CharacterPage, error) {
	var characterPages []models.CharacterPage

	url := c.http.FullPath()

	for i := 1; ; i++ {
		var characterPage = models.CharacterPage{}
		response, err := c.http.client.R().SetQueryString(fmt.Sprintf("page=%d", i)).EnableTrace().Get(url)

		if err != nil {
			log.Fatalln(fmt.Sprintf("FindAll at page %d, error: %s", i, err))
			return characterPages, err
		}

		json.Unmarshal(response.Body(), &characterPage)
		characterPages = append(characterPages, characterPage)

		log.Println(fmt.Sprintf("Success Response Character FindAll Page: %d", i))

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
