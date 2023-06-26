package unicycle

import (
	"fmt"
	"mime"
	"os"
	"strings"
)

// FetchFile simplifies the common task of downloading a file from a url, returning the file's location
// if filename does not contain an extension, an attempt will be made to infer it from the request's Content-Type header
func FetchFile(raw_url string, options FetchOptions, directory, filename string) (string, error) {
	if directory == "" {
		return "", errFetchFileNoDirectory
	}
	if filename == "" {
		return "", errFetchFileNoFilename
	}

	response, err := Fetch(raw_url, options)
	if err != nil {
		return "", err
	}

	ok, err := ResponseOk(response)
	if !ok {
		return "", err
	}

	filenameParts := strings.Split(filename, ".")
	if len(filenameParts) == 1 { // no file extension specified by caller
		contentType := response.Header.Get("Content-Type")
		if contentType == "" {
			return "", newFetchError(fmt.Errorf("response missing Content-Type header and no file extension specified in filename %v", filename), response)
		}
		extensions, err := mime.ExtensionsByType(contentType)
		if err != nil {
			return "", newFetchError(err, response)
		}
		if len(extensions) == 0 {
			return "", newFetchError(fmt.Errorf("no file extensions matched Content-Type: %v", contentType), response)
		}
		filename += extensions[0]
	}

	filePath := fmt.Sprintf("%v/%v", directory, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	defer file.Close()

	file.ReadFrom(response.Body)
	return filePath, nil
}
