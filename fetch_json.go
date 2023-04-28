package unicycle

import (
	"encoding/json"
	"io"
	"net/http"
)

// FetchJson simplifies the common task of making a HTTP request to fetch some JSON data and returning it as a struct
func FetchJson[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	ok, err := ResponseOk(response)
	if !ok {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	output, err := ResponseToJson[OUTPUT_TYPE](response)
	return output, newFetchError(err, response)
}

// Like FetchJson, but attempts to parse to the json struct regardless of status code
func FetchJsonAlways[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	output, err := ResponseToJson[OUTPUT_TYPE](response)
	return output, newFetchError(err, response)
}

func ResponseToJson[OUTPUT_TYPE any](response *http.Response) (OUTPUT_TYPE, error) {
	responseBodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	var output OUTPUT_TYPE
	err = json.Unmarshal(responseBodyBytes, &output)
	return output, err
}
