package unicycle

import (
	"encoding/json"
	"fmt"
	"io"
)

// FetchJson simplifies the common task of fetching some JSON data and returning it as a struct
func FetchJson[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return output, err
	}

	if (response.StatusCode < 200) || (300 <= response.StatusCode) {
		return output, newFetchError(fmt.Errorf("non-2XX response status code in FetchJson: %d", response.StatusCode), response)
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
