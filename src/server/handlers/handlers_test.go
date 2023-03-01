package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"github.com/gorilla/mux"
)

type mockUserService struct {
}

func (m *mockUserService) CreateUser(user *models.User) (insertedID uint64, err error) {
	return 1, nil
}

func (m *mockUserService) UpdateUser(userId uint64, user *models.User) (*models.User, error) {

	return user, nil
}

func (m *mockUserService) FindUser(userId uint64) (*models.User, error) {
	user := &models.User{ID: 1, Username: "test"}
	return user, nil
}

func (m *mockUserService) FindUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{Username: "test", Password: "$2a$14$vu94qimywaBh0McBLQ91DuNBxHuFwNiwM0x6jzAvf3wQoTXAa.w4K"}
	return user, nil
}

func (m *mockUserService) DeleteUser(userId uint64) (bool, error) {

	return true, nil
}

func TestRegisterUser(t *testing.T) {
	// create a new HTTP request with a valid user payload
	userJSON := `{"username":"test","password":"password123"}`
	req, err := http.NewRequest("POST", "/register", bytes.NewBufferString(userJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a new mock Env instance
	env := &Env{
		users: &mockUserService{},
	}

	// execute the RegisterUser function
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.RegisterUser)
	handler.ServeHTTP(rr, req)

	// check the response status code is 201 Created
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// check the response body contains the expected insertedID value
	expectedBody := "1"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestAuthenticate(t *testing.T) {
	// create a new HTTP request with valid credentials
	credentialsJSON := `{"email":"test@example.com", "user": "test", "password":"password123"}`
	req, err := http.NewRequest("POST", "/authenticate", bytes.NewBufferString(credentialsJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a new mock Env instance
	env := &Env{
		users: &mockUserService{},
	}

	// execute the Authenticate function
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.Authenticate)
	handler.ServeHTTP(rr, req)

	// check the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body contains the expected message
	expectedBody := `"User logged in."`
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestGetUser(t *testing.T) {
	env := &Env{
		users: &mockUserService{},
	}
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "1"})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(env.GetUser)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the Content-Type header is set correctly.
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

	// Check the response body is what we expect.
	expected := `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":1,"Email":"","Username":"test","Password":"","AccountType":"","FirstName":"","LastName":"","ContactInfoID":0,"BusinessID":0,"UserPermissionsID":0,"UserPreferencesID":0,"ProfilePicID":0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateUser(t *testing.T) {
	// create a new HTTP request with valid user ID and updated user data
	updatedUserJSON := `{"email":"johndoe@example.com","accounttype":"business"}`
	req, err := http.NewRequest("POST", "/user/1", bytes.NewBufferString(updatedUserJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// create a new mock Env instance
	env := &Env{
		users: &mockUserService{},
	}

	// create a new Gorilla Mux router and add the UpdateUser endpoint
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", env.UpdateUser).Methods("POST")

	// extract the user ID from the URL path and add it to the request context
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)

	// execute the UpdateUser function
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// check the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// check the response body contains the expected message
	expectedBody := `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Email":"johndoe@example.com","Username":"","Password":"","AccountType":"business","FirstName":"","LastName":"","ContactInfoID":0,"BusinessID":0,"UserPermissionsID":0,"UserPreferencesID":0,"ProfilePicID":0}`
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestDeleteUser(t *testing.T) {
	// Create a new mock environment
	env := &Env{
		users: &mockUserService{},
	}

	// Create a new request with a mock user ID
	req, err := http.NewRequest("DELETE", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// create a new Gorilla Mux router and add the UpdateUser endpoint
	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", env.DeleteUser).Methods("DELETE")

	// extract the user ID from the URL path and add it to the request context
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)

	// Create a new recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the DeleteUser function with the mock environment and request
	env.DeleteUser(rr, req)

	// Check that the response status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check that the response body is the expected JSON message
	expected := `"User deleted"`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestParseRequestID(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req = mux.SetURLVars(req, map[string]string{"id": "123"})

	id, err := parseRequestID(req)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if id != 123 {
		t.Errorf("Unexpected id value: %d", id)
	}
}
