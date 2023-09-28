package util

func ErrFirst(errs []error) error {
	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}

type Error string

func (e Error) Error() string {
	return string(e)
}
