package unicycle

import (
	"encoding/json"
	"io"
)

// FetchJson simplifies the common task of fetching some JSON data and returning it as a struct
func FetchJson[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return output, err
	}

	ok, err := ResponseOk(response)
	if !ok {
		return output, err
	}

	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output, newFetchError(err, response)
	}

	err = json.Unmarshal(responseBodyBytes, &output)
	if err != nil {
		return output, newFetchError(err, response)
	}

	return output, nil
}
