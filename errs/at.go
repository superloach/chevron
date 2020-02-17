package errs

type At struct {
	Line string
	Err  error
}

func (a At) Error() string {
	return "at line " + a.Line + ": " + a.Err.Error()
}

func (a At) Unwrap() error {
	return a.Err
}
