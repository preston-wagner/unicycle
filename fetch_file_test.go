package unicycle

import (
	"os"
	"testing"
)

func TestFetchFile(t *testing.T) {
	filepath, err := FetchFile("https://www.iana.org/_img/2022/iana-logo-header.svg", FetchOptions{}, "temp", "temp")
	if err != nil {
		t.Errorf("Error fetching test file")
	}
	if filepath != "temp/temp.svg" {
		t.Errorf("filepath was not expected value")
	}
	_, err = os.Stat(filepath)
	if err != nil {
		t.Errorf("downloaded test file not found")
	}
	os.Remove(filepath)
}
