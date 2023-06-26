package unicycle

import (
	"errors"
	"fmt"
	"net/http"
)

type FetchError struct {
	Response *http.Response
	Err      error
}

func (e FetchError) Error() string {
	return "unicycle.Fetch error: " + e.Err.Error()
}

func (e FetchError) Unwrap() error {
	return e.Err
}

type BadResponseError struct {
	StatusCode int
}

func (e BadResponseError) Error() string {
	return fmt.Sprintf("non-2XX response status code: %d", e.StatusCode)
}

var errFetchNilResponse = errors.New("response is nil")

var errFetchFileNoDirectory = errors.New("FetchFile requires a directory")
var errFetchFileNoFilename = errors.New("FetchFile requires a filename")
