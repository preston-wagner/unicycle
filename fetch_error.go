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

func newFetchError(err error, response *http.Response) error {
	if err == nil {
		return nil
	}
	return FetchError{
		Err:      err,
		Response: response,
	}
}

type BadResponseError struct {
	StatusCode int
}

func (e BadResponseError) Error() string {
	return fmt.Sprintf("non-2XX response status code: %d", e.StatusCode)
}

// a helper function to simplify checking HTTP status codes from potentially-wrapped errors
func IsBadResponseWithCode(err error, code int) bool {
	var badResponseErr BadResponseError
	if errors.As(err, &badResponseErr) {
		return badResponseErr.StatusCode == code
	}
	return false
}

var errFetchNilResponse = errors.New("response is nil")

var errFetchFileNoDirectory = errors.New("FetchFile requires a directory")
var errFetchFileNoFilename = errors.New("FetchFile requires a filename")
