package models

import (
	"strings"

	"golang.org/x/exp/slices"
)

// TODO:  Add documentation (func TrimWhitespace)
func TrimWhitespace(value string) string {
	return strings.TrimSpace(value)
}

// TODO:  Add documentation (func StandardizeUserAccountType)
func StandardizeUserAccountType(userAcctType string) string {
	return strings.ToTitle(strings.ToLower(TrimWhitespace(userAcctType)))
}

// TODO:  Add documentation (func StandardizeEmailAddress)
func StandardizeEmailAddress(emailAddress string) string {
	return strings.ToLower(TrimWhitespace(emailAddress))
}

// TODO:  Add documentation (func StandardizeEmailAddress)
func StandardizeNameField(name string) string {
	return strings.ToTitle(TrimWhitespace(name))
}

// TODO:  Add documentation (func UserAccountTypeIsValid)
func UserAccountTypeIsValid(userAcctType string) bool {
	validAccountTypes := []string{"business", "individual", "system"}

	return !slices.Contains(validAccountTypes, userAcctType)
}
