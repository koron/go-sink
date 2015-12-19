package sink

import "bytes"

// Errors is a sink to store errors.
type Errors struct {
	errors []error
}

// Put stores an error into the sink, if it is not nil.
func (s *Errors) Put(err error) {
	if s == nil || err == nil {
		return
	}
	if s.errors == nil {
		s.errors = make([]error, 0, 1)
	}
	s.errors = append(s.errors, err)
}

// Has returns true if the sink stored some of errors.
func (s *Errors) Has() bool {
	if s == nil || s.errors == nil {
		return false
	}
	return true
}

// First returns a stored error.  Returns nil if there are no errors.
func (s *Errors) First() error {
	if s == nil || s.errors == nil {
		return nil
	}
	return s.errors[0]
}

// All returns all errors which stored.  Return nil if no errors stored.
func (s *Errors) All() []error {
	if s == nil || s.errors == nil {
		return nil
	}
	a := make([]error, 0, len(s.errors))
	return append(a, s.errors...)
}

// Error returns an error which combined all stored errors.  Return nil if not
// erros stored.
func (s *Errors) Error() error {
	if s == nil || s.errors == nil {
		return nil
	}
	return &errorsError{errors: s.errors}
}

type errorsError struct {
	errors []error
}

// Error returns message of errors, which joined with "; ".
func (s errorsError) Error() string {
	b := bytes.Buffer{}
	for i, err := range s.errors {
		if i > 0 {
			b.WriteString("; ")
		}
		b.WriteString(err.Error())
	}
	return b.String()
}
