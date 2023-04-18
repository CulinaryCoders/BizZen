package tests

import (
	"server/models"
	"testing"
)

func TestHashPassword(t *testing.T) {
	var testPassword string = "password"
	var hashCost int = 8

	hashedPassword, err := models.HashPassword(testPassword, hashCost)

	// assert that the function returns no error
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// assert that the password is not empty
	if hashedPassword == "" {
		t.Errorf("expected hashed password to not be empty, but got an empty string")
	}

	// assert that the hashed password is different from the original password
	if hashedPassword == "password" {
		t.Errorf("expected hashed password to be different from the original password, but they are the same")
	}
}

func TestCheckPassword(t *testing.T) {
	user := &models.User{
		Password: "$2a$14$yYPvfXdM7SaAdxA5jQr4Bu1jq9AsqBSA4lL.8LI8FostoL1UCcth2",
	}

	// Test with correct password
	err := user.CheckPassword("password123")
	if err != nil {
		t.Errorf("CheckPassword() returned an unexpected error: %v", err)
	}

	// Test with incorrect password
	err = user.CheckPassword("wrongpassword")
	if err == nil {
		t.Error("CheckPassword() did not return an error with an incorrect password")
	}
}
