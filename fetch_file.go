package unicycle

import (
	"errors"
	"fmt"
	"mime"
	"os"
	"strings"
)

// FetchFile simplifies the common task of downloading a file from a url, returning the file's location
// if filename does not contain an extension, it will be inferred from the request's Content-Type header
func FetchFile(raw_url string, options FetchOptions, directory, filename string) (string, error) {
	if directory == "" {
		return "", errors.New("FetchFile requires a directory")
	}
	if filename == "" {
		return "", errors.New("FetchFile requires a filename")
	}
	resp, err := Fetch(raw_url, options)
	if err != nil {
		return "", err
	}
	filenameParts := strings.Split(filename, ".")
	if len(filenameParts) == 1 { // no file extension
		contentType := resp.Header.Get("Content-Type")
		if contentType == "" {
			return "", errors.New("response missing Content-Type header")
		}
		extensions, err := mime.ExtensionsByType(contentType)
		if err != nil {
			return "", err
		}
		if len(extensions) == 0 {
			return "", fmt.Errorf("no file extensions matched Content-Type: %v", contentType)
		}
		filename += extensions[0]
	}
	filePath := fmt.Sprintf("%v/%v", directory, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	defer file.Close()
	file.ReadFrom(resp.Body)
	return filePath, nil
}
