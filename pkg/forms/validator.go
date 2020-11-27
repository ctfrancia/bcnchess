package forms

import (
	"fmt"
	"regexp"
)

/*
   "clubCountry": "Spain",
   "userCountry": "Spain",
   "club": "club d'Escacs Congrés",
   "lichessUsername": "chesspanic",
   "chesscomUsername": "na"
*/
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
	model map[string]string
}

// NewValidator takes the a map of strings to validate. Returns a pointer to a new validator
func NewValidator(input map[string]string) (*Validator, error) {
	for field := range input {
		_, found := find(requiredFields, field)
		fmt.Println("found", field, found)
		if !found {
			err := fmt.Errorf("an unsupported field was sent: %s, please see docs for valid fields", field)
			return nil, err
		}
	}
	return &Validator{model: input}, nil
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

// MatchesPattern checks against our EmailRx
func (f *Validator) MatchesPattern(field string, pattern *regexp.Regexp) {
	// value := f.Get(field)
	/*
		if value == "" {
			return
		}
		if !pattern.MatchString(value) {
			f.Errors.Add(field, ErrFieldInvalid)
		}
	*/
}

// func (v *Validator)
