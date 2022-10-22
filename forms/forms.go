package forms

import (
	"net/http"
	"net/url"
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

// Checks if form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	checker := r.Form.Get(field)
	if checker == "" {
		f.Errors.Add(field, "This field cannot be empty")
	}
	return checker != ""
}
