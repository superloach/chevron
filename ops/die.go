package ops

import (
	"io"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type DIE struct {
	Text string
}

func (d DIE) String() string {
	return "DIE `" + d.Text + "`"
}

func (d DIE) Run(v *vars.Vars, _ io.Reader, w io.Writer) error {
	text, err := mix.Mix(d.Text, v)
	if err != nil {
		return err
	}

	w.Write([]byte(text))
	w.Write([]byte{'\n'})

	return errs.DIE
}
