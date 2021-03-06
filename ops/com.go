package ops

import (
	"io"

	"github.com/superloach/chevron/vars"
)

type COM struct {
	Comment string
}

func (c COM) String() string {
	return "COM `" + c.Comment + "`"
}

func (c COM) Run(v *vars.Vars, _ io.Reader, _ io.Writer) error {
	v.Set("_c", c.Comment)
	return nil
}
