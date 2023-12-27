package base

import "net/url"

type errs map[string][]string

// Form for handling form
type Form struct {
	url.Values
	Errors errs
}

// New returns the new form instance
func New(data url.Values) *Form {
	return &Form{
		data,
		make(errs),
	}
}

func (e errs) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Valid validates the form by checking if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Name gets the name entered and checks if it is empty
func (f *Form) Name(field string) *Form {
	val := f.Get(field)
	if val == "" {
		f.Errors.Add(field, "Name cannot be empty")
	}
	return f
}
