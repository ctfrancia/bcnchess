package forms

import (
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

// EmailRX defines how we check if the email is correct format
var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

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

// MinLength defines the minimum value alotted to given field
func (f *Form) MinLength(field string, d int) {
	value := f.Get(field)

	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < d {
		f.Errors.Add(field, errFieldTooShort(d))
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

// MatchesPattern checks against our EmailRx
func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, errFieldInvalid)
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
