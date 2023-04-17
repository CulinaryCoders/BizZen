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

// TODO:  Add documentation (func TrimWhitespace)
func TrimWhitespace(value string) string {
	return strings.TrimSpace(value)
}

// TODO:  Add documentation (func StandardizeEmailAddress)
func StandardizeEmailAddress(emailAddress string) string {
	return strings.ToLower(TrimWhitespace(emailAddress))
}

// TODO:  Add documentation (func StandardizeEmailAddress)
func StandardizeNameField(name string) string {
	return titleCaser.String(TrimWhitespace(name))
}

/*  --  GENERIC VALIDATION FUNCTIONS  --  */
// TODO:  Add documentation (func EmailAddressIsValid)
func EmailAddressIsValid(emailAddress string) bool {
	return true
}

/*  --  OBJECT-SPECIFIC STANDARDIZATION FUNCTIONS  --  */

/*  USER FIELDS  */
// TODO:  Add documentation (func StandardizeUserAccountType)
func StandardizeUserAccountType(userAcctType string) string {
	return titleCaser.String(strings.ToLower(TrimWhitespace(userAcctType)))
}

/*  --  OBJECT-SPECIFIC VALIDATION FUNCTIONS  --  */

/*  USER FIELDS  */
// TODO:  Add documentation (func UserAccountTypeIsValid)
func UserAccountTypeIsValid(userAcctType string) bool {
	validAccountTypes := []string{"business", "individual", "system"}

	return !slices.Contains(validAccountTypes, userAcctType)
}
