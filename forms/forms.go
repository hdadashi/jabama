package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Creates a custom form struct and embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Returns true if there are no errors, otherwise returns false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Checks if the field is required
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// Checks for string minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	value := f.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

// Checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	checker := r.Form.Get(field)
	if checker == "" {
		f.Errors.Add(field, "This field cannot be empty")
	}
	return checker != ""
}

// Checks for valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
