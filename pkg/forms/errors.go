package forms

import "fmt"

type errors map[string][]string

const (
	errCannotBeBlank = "This field cannot be blank"
	errFieldInvalid  = "This field is invalid"
)

func errFieldTooLong(d int) string {
	return fmt.Sprintf("This field is too long(maximum is %d characters)", d)
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
