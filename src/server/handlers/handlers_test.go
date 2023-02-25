package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/models"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRegisterUser(t *testing.T) {
	testDB, err := gorm.Open(postgres.Open("test DB connection string"), &gorm.Config{})
	testDB.Migrator().DropTable(&models.User{})
	testDB.AutoMigrate(&models.User{})

	h := &Handler{DB: testDB}

	user := &models.User{
		Email:    "test@example.com",
		Username: "test",
		Password: "password",
	}
	jsonBody, err := json.Marshal(user)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(jsonBody))
	if err != nil {
		t.Error(err)
	}

	responseRecorder := httptest.NewRecorder()

	h.RegisterUser(responseRecorder, req)

	if responseRecorder.Code != http.StatusCreated {
		t.Errorf("expected 201 but go %d", req.Response.StatusCode)
	}
}
