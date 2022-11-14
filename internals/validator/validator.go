// Filename: internals/validator/validator.go

package validator

import (
	"net/url"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// We create a type that wraps our validation errors map
type Validator struct {
	Errors map[string]string
}

// New() creates a new validator instance
func New() *Validator {
	return &Validator{
		Errors: make(map[string]string),
	}
}

// Valid() checks the errors map for entires
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// In() checks if an element can be found in a provided list of elements
func In(element string, list ...string) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}

// Matches() returns true if a string vaue matches a specific regex pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// ValidWebsite() checks if a string value is a valid web URl
func ValidWebsite(website string) bool {
	_, err := url.ParseRequestURI(website)
	return err == nil
}

// Adderror() adds a  erro entry yo the errors map
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check() performs the validation checks and calls the adderrors
// Method in tuen if an error entry needs to be added
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}

}

// Unique() check that there are no repeating values in the slice
func Unique(values []string) bool {
	UniqueValues := make(map[string]bool)

	for _, value := range values {
		UniqueValues[value] = true
	}
	return len(values) == len(UniqueValues)
}
