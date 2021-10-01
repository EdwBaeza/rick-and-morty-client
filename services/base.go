package services

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

const BASEURL = "https://rickandmortyapi.com/api/"

type http struct {
	resource string
	client   *resty.Client
}

// Get request to rickandmortyapi
func (h *http) get(relativePath string, pathParams map[string]string, queryParams map[string]string) ([]byte, error) {
	response, err := h.client.
		R().
		SetQueryParams(queryParams).
		SetPathParams(pathParams).
		EnableTrace().
		Get(BASEURL + relativePath)

	if err != nil || response.StatusCode() != 200 {
		return []byte{}, errors.New(response.String())
	}

	return response.Body(), nil
}
