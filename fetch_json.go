package unicycle

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FetchJsonError struct {
	Response *http.Response
	Err      error
}

func (e *FetchJsonError) Error() string {
	return e.Err.Error()
}

func NewFetchJsonError(err error, response *http.Response) *FetchJsonError {
	return &FetchJsonError{
		Err:      err,
		Response: response,
	}
}

func FetchJson[OUTPUT_TYPE any](raw_url string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	response, err := Fetch(raw_url, options)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	if (response.StatusCode < 200) || (300 <= response.StatusCode) {
		return output, NewFetchJsonError(fmt.Errorf("non-2XX response status code: %d", response.StatusCode), response)
	}
	response_body_bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	err = json.Unmarshal(response_body_bytes, &output)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	return output, nil
}
