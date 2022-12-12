package unicycle

import (
	"encoding/json"
	"fmt"
	"io"
)

func FetchJson[OUTPUT_TYPE any](raw_url string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	response, err := Fetch(raw_url, options)
	if err != nil {
		return output, err
	}
	if (response.StatusCode < 200) || (300 <= response.StatusCode) {
		return output, fmt.Errorf("non-2XX response status code: %d", response.StatusCode)
	}
	response_body_bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output, err
	}
	err = json.Unmarshal(response_body_bytes, &output)
	return output, err
}
