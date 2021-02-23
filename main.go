package main

import (
	"fmt"
	"reflect"

	"github.com/edwbaeza/rick-and-morty-client/services"
)

func main() {
	// Example FindByID Character
	characterService := services.NewCharacterService()
	character, _ := characterService.FindByID(1)
	fmt.Println("FindByID Character:")
	fmt.Println(reflect.TypeOf(character))
	fmt.Println(character)
	fmt.Println(character.ID)
	fmt.Println(character.Name)
	fmt.Println(character.Status)

	// Example FindAll Characters
	fmt.Println("FindAll Characters:")
	characterPages, _ := characterService.FindAll()
	fmt.Println(reflect.TypeOf(characterPages))
	fmt.Println("Page 1:", characterPages[0])
	fmt.Println("Page 2:", characterPages[1])
	fmt.Println("Pages:", len(characterPages))
	fmt.Println("Select location 1 from page 1:", characterPages[0].Characters[0])

	// Example FindByID Location
	locationService := services.NewLocationService()
	location, _ := locationService.FindByID(1)
	fmt.Println("FindByID Character:")
	fmt.Println(reflect.TypeOf(location))
	fmt.Println(location)
	fmt.Println(location.ID)
	fmt.Println(location.Name)

	// Example FindAll Locations
	fmt.Println("FindAll Locations:")
	locationPages, _ := locationService.FindAll()
	fmt.Println(reflect.TypeOf(locationPages))
	fmt.Println("Page 1:", locationPages[0])
	fmt.Println("Page 2:", locationPages[1])
	fmt.Println("Pages:", len(locationPages))
	fmt.Println("Select location 1 from page 1:", locationPages[0].Locations[0])
}
