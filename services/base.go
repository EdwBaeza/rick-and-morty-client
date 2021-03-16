package services

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
)

const BASEURL = "https://rickandmortyapi.com/api/"

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

func (s *http) GenerateQueryParams(params map[string]string) string {
	var fullQueryParams string

	for key, value := range params {
		format := "&%s=%s"

		if fullQueryParams == "" {
			format = "%s=%s"
		}

		fullQueryParams += fmt.Sprintf(format, key, value)
	}

	return fullQueryParams
}
