package unicycle

import (
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type FetchOptions struct {
	Method  string
	Query   map[string]string
	Headers map[string]string
	Body    io.Reader
	Timeout *time.Duration
}

type FetchError struct {
	Response *http.Response
	Err      error
}

func (e *FetchError) Error() string {
	return e.Err.Error()
}

func (e *FetchError) Unwrap() error {
	return e.Err
}

func (e *FetchError) LogDetails() {
	if e.Response == nil {
		log.Println("FetchError.LogDetails() error: no Response")
	} else {
		if e.Response.Request == nil {
			log.Println("FetchError.LogDetails() error: e.Response.Request is nil")
		} else {
			log.Println(e.Response.Request.URL)
		}
		log.Println(e.Response.StatusCode)
		log.Println(e.Response.Status)
		responseBodyBytes, err := io.ReadAll(e.Response.Body)
		if err != nil {
			log.Println("FetchError.LogDetails() error: could not read body:", err)
		} else {
			log.Println(string(responseBodyBytes))
		}
	}
}

func newFetchError(err error, response *http.Response) *FetchError {
	return &FetchError{
		Err:      err,
		Response: response,
	}
}

func LogPossibleFetchError(err error) bool {
	log.Println(err)
	var fetchError *FetchError
	if errors.As(err, &fetchError) {
		fetchError.LogDetails()
		return true
	}
	return false
}

// Fetch simplifies common http requests and associated error checking
func Fetch(raw_url string, options FetchOptions) (*http.Response, error) {
	true_url, err := url.Parse(raw_url)
	if err != nil {
		return nil, err
	}

	if len(options.Query) > 0 {
		query := true_url.Query()
		for key, value := range options.Query {
			query.Set(key, value)
		}
		true_url.RawQuery = query.Encode()
	}

	if options.Method == "" {
		options.Method = "GET"
	}

	request, err := http.NewRequest(options.Method, true_url.String(), options.Body)
	if err != nil {
		return nil, err
	}

	for key, value := range options.Headers {
		request.Header.Add(key, value)
	}

	timeout := time.Minute
	if options.Timeout != nil {
		timeout = *options.Timeout
	}

	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return response, newFetchError(err, response)
	}

	return response, nil
}
