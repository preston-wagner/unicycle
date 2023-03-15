package unicycle

import (
	"strings"
	"testing"
)

func TestFetchString(t *testing.T) {
	placeholder, err := FetchString("https://jsonplaceholder.typicode.com/todos/1", FetchOptions{})
	if err != nil {
		t.Error("Error fetching test string")
	}
	if !strings.Contains(placeholder, "delectus aut autem") {
		t.Error("FetchString result did not contain expected substring")
	}
}
