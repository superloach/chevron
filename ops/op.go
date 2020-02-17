package ops

import (
	"io"

	"github.com/superloach/chevron/vars"
)

type Op interface {
	Run(*vars.Vars, io.Reader, io.Writer) error
	String() string
}
