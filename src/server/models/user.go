package models

import (
	"log"
	"server/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// TODO: Add foreign key logic to User model
// TODO: Add constraint for AccountType column to limit user types
// GORM model for all User records in the database
type User struct {
	gorm.Model
	Email       string `gorm:"not null;unique;column:email" json:"email"`                             // User's email address
	Username    string `gorm:"not null;unique;column:username" json:"username"`                       // Username
	Password    string `gorm:"not null;column:password" json:"password"`                              // User's hashed password
	AccountType string `gorm:"not null;column:account_type" json:"account_type"`                      // Account type of the User record (Individual, Business, System)
	FirstName   string `gorm:"not null;column:first_name" json:"first_name"`                          // User's first name
	LastName    string `gorm:"not null;column:last_name" json:"last_name"`                            // User's last name
	BusinessID  *uint  `gorm:"column:business_id;default:null" json:"business_id" sql:"DEFAULT:NULL"` // ID of the Business record associated with the User record
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

func GetID

# Returns ID field from User object

*Parameters*

	N/A (None)

*Returns*

	_  <uint>

		The ID of the User object
*/
func (user *User) GetID() uint {
	return user.ID
}

/*
*Description*

func Create

Creates a new User record in the database and returns the created record along with any errors that are thrown.

*Parameters*

	db  <*gorm.DB>

		A pointer to the database instance where the record will be created.

*Returns*

	_  <*User>

		The created User record.

	_  <error>

		Encountered error (nil if no errors are encountered).
*/
func (user *User) Create(db *gorm.DB) (map[string]Model, error) {
	// TODO: Add field validation logic (func Create) -- add as BeforeCreate gorm hook definition at the top of this file
	err := db.Create(&user).Error
	returnRecords := map[string]Model{"user": user}
	return returnRecords, err
}

/*
*Description*

func Get

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
*/
func (user *User) Get(db *gorm.DB, userID uint) (map[string]Model, error) {
	err := db.First(&user, userID).Error
	returnRecords := map[string]Model{"user": user}
	return returnRecords, err
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

func Update

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
*/
func (user *User) Update(db *gorm.DB, userID uint, updates map[string]interface{}) (map[string]Model, error) {
	// Confirm userID exists in the database and get current object
	returnRecords, err := user.Get(db, userID)
	updateUser := returnRecords["user"]

	if err != nil {
		return returnRecords, err
	}

	// TODO: Add field validation logic (func Update) -- add as BeforeUpdate gorm hook definition at the top of this file
	err = db.Model(&updateUser).Where("id = ?", userID).Updates(updates).Error
	returnRecords = map[string]Model{"user": updateUser}

	return returnRecords, err
}

/*
*Description*

func Delete

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
*/
func (user *User) Delete(db *gorm.DB, userID uint) (map[string]Model, error) {
	// Confirm userID exists in the database and get current object
	returnRecords, err := user.Get(db, userID)
	deleteUser := returnRecords["user"]

	if err != nil {
		return returnRecords, err
	}

	if config.Debug {
		log.Printf("\n\nUser object targeted for deletion:\n\n%+v\n\n", deleteUser)
	}

	err = db.Delete(deleteUser).Error
	returnRecords = map[string]Model{"user": deleteUser}

	return returnRecords, err
}
