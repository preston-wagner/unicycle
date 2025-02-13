package fetch

import "github.com/nuvi/unicycle/string_ext"

// FetchString simplifies the common task of making a HTTP request to fetch some plaintext data
func FetchString(rawUrl string, options FetchOptions) (string, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return "", err
	}

	if !options.AcceptBadResponse {
		ok, err := ResponseOk(response)
		if !ok {
			return "", err
		}
	}

	body, err := string_ext.ReadString(response.Body)
	if err != nil {
		return "", newFetchError(err, response)
	}

	return body, nil
}
