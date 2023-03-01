package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add foreign key logic to User model
// TODO: Update time columns type / formatting to ensure behavior/values are expected
// TODO: Add constraint for AccountType column to limit user types
type User struct {
	gorm.Model
	Email             string `gorm:"not null;unique;column:email" json:"email"`
	Username          string `gorm:"not null;unique;column:username" json:"username"`
	Password          string `gorm:"not null;unique;column:password" json:"password"`
	AccountType       string `gorm:"not null;unique;column:account_type" json:"account_type"`
	FirstName         string `gorm:"not null;column:first_name" json:"first_name"`
	LastName          string `gorm:"not null;column:last_name" json:"last_name"`
	ContactInfoID     uint   `gorm:"column:contact_info_id" json:"contact_info_id"`
	BusinessID        uint   `gorm:"column:business_id" json:"business_id"`
	UserPermissionsID uint   `gorm:"column:permissions_id" json:"permissions_id"`
	UserPreferencesID uint   `gorm:"column:user_pref_id" json:"user_pref_id"`
	ProfilePicID      uint   `gorm:"column:profile_pic_id" json:"profile_pic_id"`
}

type UserPermissions struct {
	gorm.Model
	Label       string `gorm:"not null;column:label" json:"label"`
	Description string `gorm:"not null;column:desc" json:"desc"`
}

// Equal determines if two different User objects are equal to each other (i.e. all fields match).
//
// Parameters:
// -compareUser: The User object that the calling User object is being compared to.
//
// Returns:
// -unequalFields []string: The list of fields that did not match between the two User objects being compared
// -equal bool: If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.
//
// Description:
// This function determines if two User object instances are equal to each other. The primary purpose of this function
// is to test the functionality of database and handler calls to ensure that the correct objects are being returned and/or
// updated in the database.
func (user *User) Equal(compareUser *User) (unequalFields []string, equal bool) {
	equal = true

	if user.ID != compareUser.ID {
		equal = false
		unequalFields = append(unequalFields, "ID")
	}

	if user.Email != compareUser.Email {
		equal = false
		unequalFields = append(unequalFields, "Email")
	}

	if user.Username != compareUser.Username {
		equal = false
		unequalFields = append(unequalFields, "Username")
	}

	if user.Password != compareUser.Password {
		equal = false
		unequalFields = append(unequalFields, "Password")
	}

	if user.AccountType != compareUser.AccountType {
		equal = false
		unequalFields = append(unequalFields, "AccountType")
	}

	if user.FirstName != compareUser.FirstName {
		equal = false
		unequalFields = append(unequalFields, "FirstName")
	}

	if user.LastName != compareUser.LastName {
		equal = false
		unequalFields = append(unequalFields, "LastName")
	}

	if user.ContactInfoID != compareUser.ContactInfoID {
		equal = false
		unequalFields = append(unequalFields, "ContactInfoID")
	}

	if user.BusinessID != compareUser.BusinessID {
		equal = false
		unequalFields = append(unequalFields, "BusinessID")
	}

	if user.UserPermissionsID != compareUser.UserPermissionsID {
		equal = false
		unequalFields = append(unequalFields, "UserPermissionsID")
	}

	if user.BusinessID != compareUser.UserPreferencesID {
		equal = false
		unequalFields = append(unequalFields, "BusinessID")
	}

	if user.UserPermissionsID != compareUser.ProfilePicID {
		equal = false
		unequalFields = append(unequalFields, "UserPermissionsID")
	}

	if user.CreatedAt.Equal(compareUser.CreatedAt) {
		equal = false
		unequalFields = append(unequalFields, "CreatedAt")
	}

	if user.UpdatedAt.Equal(compareUser.UpdatedAt) {
		equal = false
		unequalFields = append(unequalFields, "UpdatedAt")
	}

	if !user.DeletedAt.Time.Equal(compareUser.DeletedAt.Time) {
		equal = false
		log.Printf("DeletedAt.Time (User):  %s\nDeletedAt.Time (compareUser):  %s", user.DeletedAt.Time, compareUser.DeletedAt.Time)
		unequalFields = append(unequalFields, "DeletedAt.Time")
	}

	if user.DeletedAt.Valid != compareUser.DeletedAt.Valid {
		equal = false
		log.Printf("DeletedAt.Valid (User):  %t\nDeletedAt.Valid (compareUser):  %t", user.DeletedAt.Valid, compareUser.DeletedAt.Valid)
		unequalFields = append(unequalFields, "DeletedAt.Valid")
	}

	return unequalFields, equal
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

/*
CreateUser creates a new user in the database and returns the inserted ID and any errors that occur.

Parameters:
- user: A pointer to a User object submitted in the request

Returns:
- The insertedID (uint64): The ID of the inserted user.
- An error object, which is nil if no error is encountered or non-nil if an error occurs while retrieving the user.
*/
func (user *User) CreateUser(db *gorm.DB) (*User, error) {
	// TODO: Add field validation logic (func CreateUser) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

/*
GetUser retrieves a user with a given ID from the database and returns a pointer to the user and an error if encountered.

Parameters:
- userId: A uint64 value representing the ID of the user to be retrieved from the database.

Returns:
- *User: A pointer to a User object representing the user with the given ID.
- error: An error object, which is nil if no error is encountered or non-nil if an error occurs while retrieving the user.
*/
func (user *User) GetUser(db *gorm.DB, userID string) (*User, error) {
	err := db.First(&user, userID).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

/*
GetUserByEmail finds a user in the database by email address and returns a pointer to the User object.

Parameters:
- userEmail: The email address of the user to find.

Returns:
- *User: A pointer to the User object representing the updated user.
- error: An error object, if any errors occurred during the search process.
*/
func (user *User) GetUserByEmail(db *gorm.DB, userEmail string) (*User, error) {
	err := db.First(&user, userEmail).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

/*
UpdateUser finds a user in the database by email address and returns a pointer to the User object.

Parameters:
- userId: A uint64 value representing the ID of the user to be retrieved from the database.
- updatedUser: A pointer to a User object containing updated information

Returns:
- *User: A pointer to the User object representing the found user.
- error: An error object, if any errors occurred during the search process.
*/
func (user *User) UpdateUser(db *gorm.DB, userID string, updates map[string]interface{}) (*User, error) {
	// Confirm user exists and get current object
	var err error
	user, err = user.GetUser(db, userID)
	if err != nil {
		return user, err
	}

	// TODO: Add field validation logic (func UpdateUser) -- add as BeforeUpdate gorm hook definition at the top of this file

	if err := db.Model(&user).Where("id = ?", userID).Updates(updates).Error; err != nil {
		return user, err
	}

	return user, nil
}

/*
DeleteUser finds a user in the database by email address and returns a pointer to the User object.

Parameters:
- userId: A uint64 value representing the ID of the user to be deleted from the database.

Returns:
- bool: Returns true if user was successfully deleted and false if otherwise.
- error: An error object, if any errors occurred during the search process.
*/
func (user *User) DeleteUser(db *gorm.DB, userID string) (*User, error) {
	// Confirm user exists and get current object
	var err error
	user, err = user.GetUser(db, userID)
	if err != nil {
		return user, err
	}

	if err := db.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
