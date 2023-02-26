package models

import (
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add foreign key logic to User model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
// TODO: Add constraint for AccountType column to limit user types
type User struct {
	gorm.Model
	ID                uint64 `gorm:"primaryKey;serial"`
	Email             string `gorm:"not null;unique"`
	Username          string `gorm:"not null;unique"`
	Password          string `gorm:"not null;unique"`
	AccountType       string `gorm:"not null;"`
	FirstName         string
	LastName          string
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

type UserEnv struct {
	DB    *gorm.DB
	Store *sessions.CookieStore
}

// CheckPassword checks if a given password matches the hashed password stored in a User struct.
//
// Parameters:
// -hashedPassword: The hashed password to be compared with the given password.
// -password: The password to be checked against the hashed password.
//
// Returns:
// -bool: If the given password matches the hashed password, true is returned. Otherwise, false is returned.
//
// Description:
// This function uses the bcrypt algorithm to compare the given password with the hashed password
// stored in a User struct. If the given password matches the hashed password, true is returned.
// Otherwise, false is returned.
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword method compares a provided password with the hashed password stored in a User struct.
//
// Parameters:
// - providedPassword: The password to be checked against the hashed password stored in the User struct.
//
// Returns:
// - error: If the provided password does not match the hashed password, an error is returned.
// Otherwise, nil is returned.
//
// Description:
// This method uses the bcrypt algorithm to compare the provided password with the hashed password
// stored in the User struct. If the provided password does not match the hashed password, an error is returned.
// Otherwise, nil is returned.
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (u *UserEnv) CreateUser(user *User) (insertedID uint64, err error) {
	result := u.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (u *UserEnv) FindUser(userId uint64) (*User, error) {
	var user User

	if err := u.DB.First(&user, User{ID: userId}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserEnv) FindUserByEmail(userEmail string) (*User, error) {
	var user User

	if err := u.DB.First(&user, User{Email: userEmail}).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserEnv) UpdateUser(userId uint64, updatedUser *User) (*User, error) {
	currentUser, err := u.FindUser(userId)
	if err != nil {
		return nil, err
	}

	if err := u.DB.Model(&currentUser).Updates(&updatedUser).Error; err != nil {
		return nil, err
	}

	return currentUser, nil
}

func (u *UserEnv) DeleteUser(userId uint64) (bool, error) {
	userToDelete, err := u.FindUser(userId)
	if err != nil {
		return false, err
	}
	if err := u.DB.Delete(userToDelete).Error; err != nil {
		return false, err
	}

	return true, nil
}
