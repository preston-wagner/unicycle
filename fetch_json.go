package unicycle

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type FetchJsonError struct {
	Response *http.Response
	Err      error
}

func (e *FetchJsonError) Error() string {
	return e.Err.Error()
}

func (e *FetchJsonError) Unwrap() error {
	return e.Err
}

func (e *FetchJsonError) LogResponseBody() {
	if e.Response == nil {
		log.Println("LogResponseBody() error: no response")
	} else {
		log.Println(e.Response.Request.URL)
		responseBodyBytes, err := io.ReadAll(e.Response.Body)
		if err != nil {
			log.Printf("LogResponseBody() error: could not read body: %s", err.Error())
		} else {
			log.Println(string(responseBodyBytes))
		}
	}
}

func NewFetchJsonError(err error, response *http.Response) *FetchJsonError {
	return &FetchJsonError{
		Err:      err,
		Response: response,
	}
}

// FetchJson simplifies the common task of fetching some JSON data and returning it as a struct
func FetchJson[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	var output OUTPUT_TYPE
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	if (response.StatusCode < 200) || (300 <= response.StatusCode) {
		return output, NewFetchJsonError(fmt.Errorf("non-2XX response status code: %d", response.StatusCode), response)
	}
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	err = json.Unmarshal(responseBodyBytes, &output)
	if err != nil {
		return output, NewFetchJsonError(err, response)
	}
	return output, nil
}
