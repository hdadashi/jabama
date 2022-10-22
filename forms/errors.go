package forms

type errors map[string][]string

// Adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Returns the first error message
func (e errors) Get(field string) string {
	error_string := e[field]
	if len(error_string) == 0 {
		return ""
	}
	return error_string[0]
}
