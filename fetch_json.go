package unicycle

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

	output, err := ReadJson[OUTPUT_TYPE](response.Body)
	return output, newFetchError(err, response)
}

// Like FetchJson, but attempts to parse to the json struct regardless of status code
func FetchJsonAlways[OUTPUT_TYPE any](rawUrl string, options FetchOptions) (OUTPUT_TYPE, error) {
	response, err := Fetch(rawUrl, options)
	if err != nil {
		return ZeroValue[OUTPUT_TYPE](), err
	}

	output, err := ReadJson[OUTPUT_TYPE](response.Body)
	return output, newFetchError(err, response)
}
