package services

import (
	"strconv"

	"github.com/go-resty/resty/v2"
)

//BASEURL API
const BASEURL = "https://rickandmortyapi.com/api/"

//HttpClient for API
type http struct {
	relativePath string
	client       *resty.Client
}

//FullPath function return full url for API
func (s *http) FullPath() string {
	return BASEURL + s.relativePath
}

//FullPathWithID function return full url with path param for API
func (s *http) FullPathWithID(id int) string {
	return s.FullPath() + "/" + strconv.Itoa(id)
}
