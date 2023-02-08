package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add comment documentation (type User)
type User struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey;serial"`
	Name        string `json:"name"`
	Username    string `json:"username" gorm:"unique"`
	Email       string `json:"email" gorm:"unique"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	AccountType string `json:"account_type"`
}

// TODO: Add comment documentation (func HashPassword)
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// TODO: Add comment documentation (func CheckPassword)
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
