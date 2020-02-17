package errs

type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	EOF Err = "eof"
	DIE Err = "die"
)
