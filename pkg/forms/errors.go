package forms

import "fmt"

type errors map[string][]string

const (
	// ErrCannotBeBlank is a constant "This field cannot be blank"
	ErrCannotBeBlank = "This field cannot be blank"
	// ErrFieldInvalid is a constant "This field is invalid"
	ErrFieldInvalid = "This field is invalid"
)

// ErrFieldTooLong returns a string with too long of argument
func ErrFieldTooLong(d int) string {
	return fmt.Sprintf("This field is too long(maximum is %d characters)", d)
}

// ErrFieldTooShort returns a string with too short of argument
func ErrFieldTooShort(d int) string {
	return fmt.Sprintf("This field is too short(minumum is %d characters)", d)
}

// Add will add an error to the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// GET SPECIFIC ERROR
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
