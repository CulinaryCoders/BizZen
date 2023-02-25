package models

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	user := &User{Password: "password"}

	err := user.HashPassword(user.Password)

	// assert that the function returns no error
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	// assert that the password is not empty
	if user.Password == "" {
		t.Errorf("expected hashed password to not be empty, but got an empty string")
	}

	// assert that the hashed password is different from the original password
	if user.Password == "password" {
		t.Errorf("expected hashed password to be different from the original password, but they are the same")
	}
}

func TestCheckPassword(t *testing.T) {
	user := &User{Password: "password"}

	// Test with correct password
	err := user.CheckPassword("password")
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test with incorrect password
	err = user.CheckPassword("wrongpassword")
	if err == nil {
		t.Error("Expected an error, but got none")
	}
}
