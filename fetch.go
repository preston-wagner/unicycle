package unicycle

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

type FetchOptions struct {
	Method  string
	Query   map[string]string
	Headers map[string]string
	Body    io.Reader
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
	client := http.Client{
		Timeout: time.Minute,
	}
	return client.Do(request)
}
