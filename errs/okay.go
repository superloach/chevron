package errs

func Okay(err error) bool {
	return err == nil || err == EOF || err == DIE
}
