## Rick and Morty Client (Wrapper for [Rick and Morty API](https://rickandmortyapi.com/api/))

```go
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
	fmt.Println("Pages:", len(characterPages))
	fmt.Println("Select character 1 from page 1:", characterPages[0].Characters[0])

	// Example Filter Characters
	// Available keys
	// name: filter by the given name.
	// status: filter by the given status (alive, dead or unknown).
	// species: filter by the given species.
	// type: filter by the given type.
	// gender: filter by the given gender (female, male, genderless or unknown).

	queryParams := map[string]string{
		"name":   "rick",
		"status": "alive",
	}

	fmt.Println("Filter Characters:")
	filteredCharacterPages, _ := characterService.Filter(queryParams)
	fmt.Println(reflect.TypeOf(filteredCharacterPages))
	fmt.Println("Page 1:", filteredCharacterPages[0])
	fmt.Println("Pages:", len(filteredCharacterPages))
	fmt.Println("Select character 1 from page 1:", filteredCharacterPages[0].Characters[0])

	// Example FindByID Location
	locationService := services.NewLocationService()
	location, _ := locationService.FindByID(1)
	fmt.Println("FindByID Location:")
	fmt.Println(reflect.TypeOf(location))
	fmt.Println(location)
	fmt.Println(location.ID)
	fmt.Println(location.Name)

	// Example FindAll Locations
	fmt.Println("FindAll Locations:")
	locationPages, _ := locationService.FindAll()
	fmt.Println(reflect.TypeOf(locationPages))
	fmt.Println("Page 1:", locationPages[0])
	fmt.Println("Pages:", len(locationPages))
	fmt.Println("Select location 1 from page 1:", locationPages[0].Locations[0])

	// Example Filter Locations
	// Available keys
	// name: filter by the given name.
	// type: filter by the given type.
	// dimension: filter by the given dimension.

	queryParamsLocation := map[string]string{
		"name": "earth",
		"type": "planet",
	}

	fmt.Println("Filter Location:")
	filteredLocationPages, _ := locationService.Filter(queryParamsLocation)
	fmt.Println(reflect.TypeOf(filteredLocationPages))
	fmt.Println("Page 1:", filteredLocationPages[0])
	fmt.Println("Pages:", len(filteredLocationPages))
	fmt.Println("Select location 1 from page 1:", filteredLocationPages[0].Locations[0])

```
