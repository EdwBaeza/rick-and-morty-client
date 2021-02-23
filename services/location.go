package services

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/edwbaeza/rick-and-morty-client/models"
	"github.com/go-resty/resty/v2"
)

// Location wrap for api
type Location struct {
	http
}

// FindByID returns single Location
func (loc *Location) FindByID(id int) (models.Location, error) {
	location := models.Location{}
	response, err := loc.http.client.R().
		EnableTrace().
		Get(loc.http.FullPathWithID(id))

	if err != nil {
		log.Fatalln(fmt.Sprintf("FindByID %d, Error %s", id, err))
		return location, err
	}

	log.Println(fmt.Sprintf("Success Response Location FindByID: %d", id))
	json.Unmarshal(response.Body(), &location)
	return location, nil
}

// FindAll returns all locations by paging API
func (loc *Location) FindAll() ([]models.LocationPage, error) {
	var locationPages []models.LocationPage

	for i := 1; ; i++ {
		var locationPage = models.LocationPage{}
		response, err := loc.http.client.R().
			SetQueryString(fmt.Sprintf("page=%d", i)).
			EnableTrace().
			Get(loc.http.FullPath())

		if err != nil {
			log.Fatalln(fmt.Sprintf("FindAll at page %d, error: %s", i, err))
			return locationPages, err
		}

		json.Unmarshal(response.Body(), &locationPage)
		locationPages = append(locationPages, locationPage)

		log.Println(fmt.Sprintf("Success Response Location FindAll Page: %d", i))

		if locationPage.Info.Next == "" {
			log.Println("Pagination finished")
			break
		}
	}

	return locationPages, nil
}

// NewLocationService return new object of location service
func NewLocationService() *Location {
	return &Location{
		http: http{
			relativePath: "location",
			client:       resty.New(),
		},
	}
}
