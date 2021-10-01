package services

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/edwbaeza/rick-and-morty-client/structs"
	"github.com/go-resty/resty/v2"
)

type Location struct {
	http
}

// FindByID returns single Location
func (loc *Location) FindByID(id int) (structs.Location, error) {
	location := structs.Location{}
	pathParams := map[string]string{"id": strconv.Itoa(id)}
	queryParams := map[string]string{}
	response, err := loc.get(loc.resource+"/{id}", pathParams, queryParams)

	if err != nil {
		log.Fatalln(fmt.Sprintf("FindByID %d, Error %s", id, err))
		return location, err
	}

	log.Println(fmt.Sprintf("Success Response FindByID Location: %d", id))
	json.Unmarshal(response, &location)
	return location, nil
}

// FindAll returns all locations by paging API
func (loc *Location) FindAll() ([]structs.LocationPage, error) {
	var locationPages []structs.LocationPage
	pathParams := map[string]string{}
	queryParams := map[string]string{}

	for i := 1; ; i++ {
		locationPage := structs.LocationPage{}
		queryParams["page"] = strconv.Itoa(i)
		response, err := loc.get(loc.resource, pathParams, queryParams)

		if err != nil {
			log.Fatalln(fmt.Sprintf("FindAll at page %d, error: %s", i, err))
			return locationPages, err
		}

		json.Unmarshal(response, &locationPage)
		locationPages = append(locationPages, locationPage)

		if locationPage.Info.Next == "" {
			break
		}
	}

	return locationPages, nil
}

// Filter locations by following keys
// name: filter by the given name.
// type: filter by the given type.
// dimension: filter by the given dimension.
func (loc *Location) Filter(params map[string]string) ([]structs.LocationPage, error) {
	locationPages := []structs.LocationPage{}
	pathParams := map[string]string{}

	for i := 1; ; i++ {
		locationPage := structs.LocationPage{}
		params["page"] = strconv.Itoa(i)
		response, err := loc.get(loc.resource, pathParams, params)

		if err != nil {
			log.Fatalln(fmt.Sprintf("Filter at page %d, error: %s", i, err))
			return locationPages, err
		}

		json.Unmarshal(response, &locationPage)
		locationPages = append(locationPages, locationPage)

		log.Println(fmt.Sprintf("Success Response Filter Location Page: %d", i))

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
			resource: "location",
			client:   resty.New(),
		},
	}
}
