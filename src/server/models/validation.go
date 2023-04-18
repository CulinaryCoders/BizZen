package models

import (
	"strings"

	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*  --  GLOBAL DEFINITIONS  --  */
var titleCaser cases.Caser = cases.Title(language.English)

/*  --  GENERIC STANDARDIZATION FUNCTIONS  --  */

/*
*Description*

func TrimWhitespace

Removes leading and trailing whitespace from string.

*Parameters*

	    value <string>

			The original string value to be trimmed.

*Returns*

	    _ <string>

			The updated string value with leading/trailing whitespace removed.
*/
func TrimWhitespace(value string) string {
	return strings.TrimSpace(value)
}

/*
*Description*

func StandardizeEmailAddress

Standardizes an email address string by converting it to lower case and removing trailing/leading whitespace.

*Parameters*

	    emailAddress <string>

			The email address string to be standardized.

*Returns*

	    _ <string>

			The updated string value that has been standardized.
*/
func StandardizeEmailAddress(emailAddress string) string {
	return strings.ToLower(TrimWhitespace(emailAddress))
}

/*
*Description*

func StandardizeNameField

Standardizes a name string by converting it to title case and removing trailing/leading whitespace.

*Parameters*

	    name <string>

			The name string to be standardized.

*Returns*

	    _ <string>

			The updated string value that has been standardized.
*/
func StandardizeNameField(name string) string {
	return titleCaser.String(TrimWhitespace(name))
}

/*  --  GENERIC VALIDATION FUNCTIONS  --  */

/*
*Description*

func EmailAddressIsValid

Checks if an email address string has a valid format.

*Parameters*

	    emailAddress <string>

			The email address to be validated

*Returns*

	    _ <bool>

			'true' if email address format is valid, else 'false'.
*/
func EmailAddressIsValid(emailAddress string) bool {
	// TODO:  Implement function (EmailAddressIsValid)
	return true
}

/*  --  OBJECT-SPECIFIC STANDARDIZATION FUNCTIONS  --  */

/*  USER FIELDS  */

/*
*Description*

func StandardizeUserAccountType

Standardizes the string formatting for a User's account type by converting it to title case and removing trailing/leading whitespace.

*Parameters*

	    userAcctType <string>

			The account type string to be standardized.

*Returns*

	    _ <string>

			The updated string value that has been standardized.
*/
func StandardizeUserAccountType(userAcctType string) string {
	return titleCaser.String(strings.ToLower(TrimWhitespace(userAcctType)))
}

/*  --  OBJECT-SPECIFIC VALIDATION FUNCTIONS  --  */

/*
*Description*

func UserAccountTypeIsValid

Checks if User's account type is a valid value.

Valid account types:
  - User
  - Business
  - System

*Parameters*

	    userAcctType <string>

			The account type to be validated.

*Returns*

	    _ <bool>

			'true' if account type is a valid value, else 'false'.
*/
func UserAccountTypeIsValid(userAcctType string) bool {
	validAccountTypes := []string{"business", "user", "system"}
	stdUserAcctType := StandardizeUserAccountType(userAcctType)
	return slices.Contains(validAccountTypes, stdUserAcctType)
}
