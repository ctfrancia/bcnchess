package forms

import (
	"net/url"
	"strings"
	"unicode/utf8"
)

// Form describes our form structure
type Form struct {
	url.Values
	Errors errors
}

// New returns a pointer to the Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required is a helper for setting up our error field
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, errCannotBeBlank)
		}
	}
}

// MaxLength deines the maximum length for a given field
func (f *Form) MaxLength(field string, d int) {
	value := f.Get(field)

	if value == "" {
		return
	}

	if utf8.RuneCountInString(value) > d {
		f.Errors.Add(field, errFieldTooLong(d))
	}
}

// PermittedValues loops over the field and the options provided if there is an error then add it
func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)

	if value == "" {
		return
	}

	for _, opt := range opts {
		if value == opt {
			return
		}
	}

	f.Errors.Add(field, errFieldInvalid)
}

// Valid returns a boolean based on if there are errors within the map
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
