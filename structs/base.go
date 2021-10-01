package structs

import (
	"time"
)

//Info metadata
type Info struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type Base struct {
	ID      int       `json:"id"`
	Created time.Time `json:"created"`
}
