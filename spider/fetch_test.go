package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHTMLBody(t *testing.T) {
	const OK = "OK"
	// create a new server
	server := httptest.NewServer(
		// create a handler
		http.HandlerFunc(
			// callback for each request
			func(rw http.ResponseWriter, req *http.Request) {
				// Test case: Send response to be returned.
				rw.Write([]byte(OK))
			}))
	// Close the server when the test finishes
	defer server.Close()

	// Test the function
	bodyText, err := getHTMLBody(server.URL) // url is like localhost:1234
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	// Read from the body and check if the response matches the one defined in the test server
	if bodyText != OK {
		t.Errorf("Expected body to contain \"%s\", got: %s", OK, bodyText)
	}
}

func TestGetHTMLBodyError(t *testing.T) {
	// Here, we're using an invalid URL to simulate an error.
	body, err := getHTMLBody("http://invalidurl")
	if err == nil {
		t.Error("Expected an error")
	}
	if body != "" {
		t.Error("Expected no body")
	}
}
