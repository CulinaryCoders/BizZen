package models

import (
	"golang.org/x/crypto/bcrypt"
)

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
func HashPassword(password string, hashCost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)
	if err != nil {
		return password, err
	}

	return string(bytes), nil
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
