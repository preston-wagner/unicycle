package unicycle

import (
	"io"
)

// FetchString simplifies the common task of making a HTTP request to fetch some plaintext data
func FetchString(rawUrl string, options FetchOptions) (string, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return "", err
	}

	ok, err := ResponseOk(response)
	if !ok {
		return "", err
	}

	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", newFetchError(err, response)
	}

	return string(responseBodyBytes), nil
}
