package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"
)

type mockUserService struct {
	UsersRegistered []*models.User
}

//TODO: Refactor and implmenent mock DB functions from models.User

func (m *mockUserService) CreateUser(user *models.User) (insertedID uint64, err error) {
	return 0, nil
}

func (m *mockUserService) UpdateUser(userId uint64, user *models.User) (*models.User, error) {
	return user, nil
}

func (m *mockUserService) FindUser(userId uint64) (*models.User, error) {
	user := &models.User{Username: "test", Password: "test"}
	return user, nil
}

func (m *mockUserService) FindUserByEmail(userEmail string) (*models.User, error) {
	user := &models.User{Username: "test", Password: "test"}
	return user, nil
}

func (m *mockUserService) DeleteUser(userId uint64) (bool, error) {

	return true, nil
}

func TestRegisterUser(t *testing.T) {

	//TODO: Create working test for Register function
	t.Run("can register valid users", func(t *testing.T) {
		var jsonData = []byte(`{
			"username": "test",
			"password": "123"
		}`)

		responseRecorder := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))

		mockEnv := Env{users: &mockUserService{}}

		http.HandlerFunc(mockEnv.RegisterUser).ServeHTTP(responseRecorder, req)

		if responseRecorder.Code != http.StatusCreated {
			t.Errorf("expected 201 but go %d", responseRecorder.Code)
		}
	})

}
