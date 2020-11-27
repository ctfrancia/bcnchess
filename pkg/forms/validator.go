package forms

import (
	"fmt"
	"regexp"
)

var requiredFields = []string{
	"firstName",
	"lastName",
	"email",
	"password",
	"clubCountry",
	"userCountry",
	"clubName",
	"lichessUsername",
	"chesscomUsername",
}

// EmailRX defines how we check if the email is correct format
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator defines how validator
type Validator struct {
	model  map[string]string
	errors map[string][]string
	fields []string
}

// NewValidator takes the a map of strings to validate. Returns a pointer to a new validator
func NewValidator() *Validator {
	return &Validator{model: map[string]string{}, errors: map[string][]string{}, fields: requiredFields}
}

// ValidateFields handles the map of the request and checks against the requiredFields
func (v *Validator) ValidateFields(input map[string]string) {
	for field, value := range input {
		_, found := find(requiredFields, field)
		if !found {
			v.errors["errors"] = append(v.errors[field], fmt.Sprintf("%+v is not valid, check docs for valid fields", field))
		}
		(v.model)[field] = value
	}
}

// MatchesPattern checks against our EmailRx
func (v *Validator) MatchesPattern(field string, pattern *regexp.Regexp) {
	_, found := find(v.fields, field)
	if !found {
		v.errors["errors"] = append(v.errors["errors"], fmt.Sprintf("field not found: %s", field))
	}

	if !pattern.MatchString(v.model[field]) {
		v.errors["errors"] = append(v.errors["errors"], "email is not valid")
	}
}

// Valid returns a boolean based on if there are errors within the errors field
// func (v *Validator) Valid() bool {
// func (v *Validator) Valid() (bool, map[string][]string) {
func (v *Validator) Valid() (bool, map[string][]string) {
	if len(v.errors) == 0 {
		return true, nil
	}
	return false, v.errors
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
