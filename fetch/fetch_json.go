package fetch

import (
	"log"

	"github.com/nuvi/unicycle/defaults"
)

// FetchJson simplifies the common task of making a HTTP request to fetch some JSON data and returning it as a struct
func FetchJson[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return defaults.ZeroValue[OUTPUT_TYPE](), err
	}

	if !options.AcceptBadResponse {
		ok, err := ResponseOk(response)
		if !ok {
			return defaults.ZeroValue[OUTPUT_TYPE](), err
		}
	}

	body, err := ReadString(response.Body)
	if err != nil {
		return defaults.ZeroValue[OUTPUT_TYPE](), newFetchError(err, response)
	}

	if options.Logging {
		log.Println(body)
	}

	output, err := ReadJsonString[OUTPUT_TYPE](body)
	return output, newFetchError(err, response)
}
