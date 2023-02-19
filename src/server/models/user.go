package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add foreign key logic to User model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
// TODO: Add constraint for AccountType column to limit user types
type User struct {
	gorm.Model
	ID                uint   `gorm:"primaryKey;serial"`
	Email             string `gorm:"not null;unique"`
	Username          string `gorm:"not null;unique"`
	Password          string `gorm:"not null;unique"`
	AccountType       string `gorm:"not null;unique"`
	FirstName         string `gorm:"not null"`
	LastName          string `gorm:"not null"`
	ContactInfoID     uint
	BusinessID        uint
	UserPermissionsID uint
	UserPreferencesID uint
	ProfilePicID      uint
}

type UserPermissions struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey;serial"`
	Label       string `gorm:"not null"`
	Description string `gorm:"not null"`
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
