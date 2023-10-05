package fetch

import (
	"errors"
	"os"
	"testing"

	"github.com/preston-wagner/unicycle/error_ext"
)

func testTempFolderIsEmpty(t *testing.T) {
	files, err := os.ReadDir("temp/")
	if err != nil {
		t.Error("error reading contents of temp/")
	}
	if len(files) > 0 {
		t.Error("temp/ is not empty")
	}
}

const tempFolderPath = "temp"

func TestFetchFile(t *testing.T) {
	os.Mkdir(tempFolderPath, os.ModePerm) // ignore error if folder already exists
	filepath, err := FetchFile("https://www.iana.org/_img/2022/iana-logo-header.svg", FetchOptions{}, tempFolderPath, "temp")
	if err != nil {
		t.Error("Error fetching test file", err)
	}
	if filepath != "temp/temp.svg" {
		t.Error("filepath was not expected value")
	}
	_, err = os.Stat(filepath)
	if err != nil {
		t.Error("downloaded test file not found")
	}
	os.Remove(filepath)
	testTempFolderIsEmpty(t)
}

func TestFetchFileWith404(t *testing.T) {
	os.Mkdir(tempFolderPath, os.ModePerm) // ignore error if folder already exists
	filepath, err := FetchFile("https://www.google.com/badUrl", FetchOptions{}, tempFolderPath, "temp")
	if filepath != "" {
		t.Error("404 response did not return empty string")
	}
	if err == nil {
		t.Error("404 response did not return error")
	}
	if !IsBadResponseWithCode(err, 404) {
		t.Error("404 response did not return error wrapping BadResponseError")
	}
	if fetchError := error_ext.ErrorAs[FetchError](err); fetchError == nil {
		t.Error("FetchFile should have responded with an instance of FetchError")
	}
	var fetchError FetchError
	if !errors.As(err, &fetchError) {
		t.Error("FetchFile should have responded with an instance of FetchError (according to errors.As)")
	}
	testTempFolderIsEmpty(t)
}
