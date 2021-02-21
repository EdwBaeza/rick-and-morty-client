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
fmt.Println("Page 2:", characterPages[1])
fmt.Println("Pages:", len(characterPages))
fmt.Println("Select Character 1 from page 1:", characterPages[0].Characters[0])
```
