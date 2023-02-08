package unicycle

import (
	"testing"
)

func TestFetch(t *testing.T) {
	response, err := Fetch("https://www.google.com/", FetchOptions{})
	if err != nil {
		t.Error("Error returned fetching google homepage")
	}
	if response == nil {
		t.Error("Missing response fetching google homepage")
	}
}

func TestFetch404(t *testing.T) {
	response, err := Fetch("https://www.google.com/badUrl", FetchOptions{})
	if err != nil {
		t.Error("Fetch should not return an error on 404")
	}
	if response == nil {
		t.Error("Missing response fetching google 404")
	}
}

func TestFetchBadDomain(t *testing.T) {
	response, err := Fetch("https://www.thiscantpossiblybearealwebsite.com/", FetchOptions{})
	if err == nil {
		t.Error("Fetch should return an error if the domain doesn't exist")
	}
	if response != nil {
		t.Error("Fetch should not return a response if the domain doesn't exist")
	}
}

func TestAppendQueryParams(t *testing.T) {
	base := "https://www.google.com/search?q=lorem"
	appended, err := AppendQueryParams(base, map[string]string{
		"ie": "UTF-8",
	})
	if err != nil {
		t.Error("AppendQueryParams should not return an error if the base url was fine")
	}
	okResults := []string{
		"https://www.google.com/search?q=lorem&ie=UTF-8",
		"https://www.google.com/search?ie=UTF-8&q=lorem",
	}
	if !Includes(okResults, appended) {
		t.Errorf("AppendQueryParams failed: %v", appended)
	}
}
