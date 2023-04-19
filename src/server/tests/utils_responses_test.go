package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"server/utils"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

/*
*Description*

func TestRespondWithJSON

Tests the RespondWithJSON method and ensures that the response being returned by the method is formatted correctly and returns what is expected.
*/
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

/*
*Description*

func TestRespondWithError

Tests the RespondWithError method and ensures that the response being returned by the method is formatted correctly and returns what is expected
*/
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

/*
*Description*

func TestParseRequestID

Tests the ParseRequestID method to confirm that the ID field from the request URL is parsed into uint format and that the appropriate error is returned if the ID is missing or formatted incorrectly.
*/
func TestParseRequestID(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Failed to create request: %s", err)
	}

	var testUserID uint = 123
	idKey := "id"

	requestVars := map[string]string{
		idKey: strconv.FormatUint(uint64(testUserID), 10),
	}

	req = mux.SetURLVars(req, requestVars)

	returnedID, err := utils.ParseRequestID(req)
	if err != nil {
		t.Errorf("Encountered unexpected error when retrieving 'id' variable from request URL: %s", err)
	}

	assert.Equal(t, returnedID, testUserID, "Returned ID value (%d) should match expected ID value (%d).", returnedID, testUserID)
}

/*
*Description*

func TestParseRequestIDField

Tests the ParseRequestIDField method to confirm that the specified ID field from the request URL is parsed into uint format and that the appropriate error is returned if the field is missing or formatted incorrectly.
*/
func TestParseRequestIDField(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("Failed to create request: %s", err)
	}

	svcIDFieldKey := "service-id"
	apptIDFieldKey := "appointment-id"

	var testSvcID uint = 123
	var testApptID uint = 456

	requestVars := map[string]string{
		svcIDFieldKey:  strconv.FormatUint(uint64(testSvcID), 10),
		apptIDFieldKey: strconv.FormatUint(uint64(testApptID), 10),
	}

	req = mux.SetURLVars(req, requestVars)

	returnedSvcID, err := utils.ParseRequestIDField(req, svcIDFieldKey)
	if err != nil {
		t.Errorf("Unexpected error when getting '%s' variable from request.  --  %s", svcIDFieldKey, err)
	}

	returnedApptID, err := utils.ParseRequestIDField(req, apptIDFieldKey)
	if err != nil {
		t.Errorf("Unexpected error when getting '%s' variable from request.  --  %s", apptIDFieldKey, err)
	}

	assert.Equal(t, returnedSvcID, testSvcID, "Returned '%s' ID value (%d) should match expected ID value (%d).", svcIDFieldKey, returnedSvcID, testSvcID)
	assert.Equal(t, returnedApptID, testApptID, "Returned '%s' ID value (%d) should match expected ID value (%d).", apptIDFieldKey, returnedApptID, testApptID)

}
