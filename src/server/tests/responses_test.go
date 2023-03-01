package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"server/utils"

	"github.com/gorilla/mux"
)

// TODO:  Add documentation (func TestRespondWithJSON)
func TestRespondWithJSON(t *testing.T) {
	// Create a mock HTTP response writer
	w := httptest.NewRecorder()

	// Create a test payload
	payload := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "John",
		Age:  30,
	}

	// Call the RespondWithJSON function with the mock writer, HTTP status code, and test payload
	utils.RespondWithJSON(w, http.StatusOK, payload)

	// Check if the Content-Type header is set to "application/json"
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type header is not application/json")
	}

	// Check if the HTTP status code is correct
	if w.Code != http.StatusOK {
		t.Errorf("Expected HTTP status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check if the response body matches the expected JSON string
	expected := `{"name":"John","age":30}`
	if w.Body.String() != expected {
		t.Errorf("Expected response body %q, but got %q", expected, w.Body.String())
	}
}

// TODO:  Add documentation (func TestRespondWithError)
func TestRespondWithError(t *testing.T) {
	// Create a mock HTTP response
	w := httptest.NewRecorder()

	// Call RespondWithError with a mock error message
	utils.RespondWithError(w, http.StatusInternalServerError, "Oops! Something went wrong.")

	// Check that the response has the correct status code
	if w.Code != http.StatusInternalServerError {
		t.Errorf("Unexpected status code: got %v, expected %v", w.Code, http.StatusInternalServerError)
	}

	// Check that the response has the correct content type
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Unexpected content type: got %v, expected application/json", w.Header().Get("Content-Type"))
	}

	// Check that the response body is a valid JSON object with an "error" field
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to unmarshal response body as JSON: %v", err)
	}
	if _, ok := response["error"]; !ok {
		t.Errorf("Response body is missing expected 'error' field")
	}
}

// TODO:  Add documentation (func TestParseRequestID)
func TestParseRequestID(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	id, err := utils.ParseRequestID(req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != 123 {
		t.Errorf("Unexpected id value: %d", id)
	}
}
