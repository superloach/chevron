package ops

import (
	"io"

	"github.com/superloach/chevron/vars"
)

type EMP struct{}

func (e EMP) String() string {
	return "EMP"
}

func (e EMP) Run(v *vars.Vars, _ io.Reader, _ io.Writer) error {
	return nil
}
