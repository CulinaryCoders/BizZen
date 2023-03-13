package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add foreign key logic to User model
// TODO: Add constraint for AccountType column to limit user types
// GORM model for all User records in the database
type User struct {
	gorm.Model
	Email             string `gorm:"not null;unique;column:email" json:"email"`                             // User's email address
	Username          string `gorm:"not null;unique;column:username" json:"username"`                       // Username
	Password          string `gorm:"not null;column:password" json:"password"`                              // User's hashed password
	AccountType       string `gorm:"not null;column:account_type" json:"account_type"`                      // Account type of the User record (Individual, Business, System)
	FirstName         string `gorm:"not null;column:first_name" json:"first_name"`                          // User's first name
	LastName          string `gorm:"not null;column:last_name" json:"last_name"`                            // User's last name
	ContactInfoID     *uint  `gorm:"column:contact_info_id" json:"contact_info_id" sql:"DEFAULT:NULL"`      // ID of ContactInfo record associated with the User record
	BusinessID        *uint  `gorm:"column:business_id;default:null" json:"business_id" sql:"DEFAULT:NULL"` // ID of the Business record associated with the User record
	UserPermissionsID *uint  `gorm:"column:permissions_id" json:"permissions_id" sql:"DEFAULT:NULL"`        // ID of the UserPermissions record associated with the User
	UserPreferencesID *uint  `gorm:"column:user_pref_id" json:"user_pref_id" sql:"DEFAULT:NULL"`            // ID of the UserPreferences record associated with the User record
	ProfilePicID      *uint  `gorm:"column:profile_pic_id" json:"profile_pic_id" sql:"DEFAULT:NULL"`        // ID of the ProfilePic record associated with the User record
}

// GORM model for all UserPermissions records in the database
type UserPermissions struct {
	gorm.Model
	Label       string `gorm:"not null;column:label" json:"label"` // Label / shortname for the permissions set
	Description string `gorm:"not null;column:desc" json:"desc"`   // Description of permissions
}

/*
*Description*

func Equal

Determines if two different User objects are equal to each other (i.e. all fields match).

The primary purpose of this function is to test the functionality of database and handler calls to ensure that
the correct objects are being returned and/or updated in the database.

*Parameters*

	compareUser  <*User>

		The User object that the calling User object is being compared to

*Returns*

	unequalFields  <[]string>

		The list of fields that did not match between the two User objects being compared

	equal  <bool>

		If all the fields between the two objects are the same, true is returned. Otherwise, false is returned.

*Response format*

	N/A (None)
*/
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

/*
*Description*

func HashPassword

Generates a hash from the provided password string and assigns it to the calling User's Password attribute.

This conforms to best practice of storing hashed passwords in the application database, rather than plain text.

*Parameters*

	password  <string>

		The plain text password that will be hashed.

*Returns*

	_  <error>

		Encountered error (nil if no errors encountered).

*Response format*

	N/A (None)
*/
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

/*
*Description*

func CheckPassword

Checks if a given password matches the hashed password associated with the calling User record's account.

This function uses the bcrypt algorithm to compare the given password with the hashed password stored in the calling User struct.

If the given password matches the hashed password, nil is returned.

*Parameters*

	password  <string>

		The password to be checked against the calling User's hashed password.

*Returns*

	_  <error>

		Encountered error (nil if no errors are encountered)

*Response format*

	N/A (None)
*/
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

/*
*Description*

func CreateUser

Creates a new User record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*User>

		The created User record.

	_  <error>

		Encountered error (nil if no errors are encountered).

*Response format*

	N/A (None)
*/
func (user *User) CreateUser(db *gorm.DB) (*User, error) {
	// TODO: Add field validation logic (func CreateUser) -- add as BeforeCreate gorm hook definition at the top of this file
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

/*
*Description*

func GetUser

Retrieves a User record in the database by ID if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	userID <uint>

		The ID of the User record being requested.

*Returns*

	_  <*User>

		The User record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)

*Response format*

	N/A (None)
*/
func (user *User) GetUser(db *gorm.DB, userID uint) (*User, error) {
	err := db.First(&user, userID).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

/*
*Description*

func GetUserByEmail

Retrieves a User record in the database by email if it exists and returns that record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve the specified record.

	userEmail  <string>

		The email of the User record being requested.

*Returns*

	_  <*User>

		The User record that is retrieved from the database.

	_  <error>

		Encountered error (nil if no errors are encountered)

*Response format*

	N/A (None)
*/
func (user *User) GetUserByEmail(db *gorm.DB, userEmail string) (*User, error) {
	err := db.First(&user, userEmail).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

/*
*Description*

func UpdateUser

Updates the specified User record in the database with the specified changes if the record exists.

Returns the updated record along with any errors that are thrown.

This function behaves like a PATCH method, rather than a true PUT. Any fields that aren't specified in the request body for the PUT request will not be altered for the specified record.

If a specified field's value should be deleted from the record, the appropriate null/blank should be specified for that key in the JSON request body (e.g. "type": "").

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance that will be used to retrieve and update the specified record.

	userID  <uint>

		The ID of the User record being updated.

	updates  <map[string]interface{}>

		JSON with the fields that will be updated as keys and the updated values as values.

		Ex:
			{
				"name": "New name",
				"address": "New address"
			}

*Returns*

	_  <*User>

		The User record that is updated in the database.

	_  <error>

		Encountered error (nil if no errors are encountered)

*Response format*

	N/A (None)
*/
func (user *User) UpdateUser(db *gorm.DB, userID uint, updates map[string]interface{}) (*User, error) {
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
*Description*

func DeleteUser

Deletes the specified User record from the database by ID if it exists.

Deleted record is returned along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be deleted from.

	userID  <uint>

		The ID of the User record being deleted.

*Returns*

	_  <*User>

		The deleted User record.

	_  <error>

		Encountered error (nil if no errors are encountered).

*Response format*

	N/A (None)
*/
func (user *User) DeleteUser(db *gorm.DB, userID uint) (*User, error) {
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
