package ops

import (
	"io"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type OUT struct {
	Text string
}

func (o OUT) String() string {
	return "OUT `" + o.Text + "`"
}

func (o OUT) Run(v *vars.Vars, _ io.Reader, w io.Writer) error {
	text, err := mix.Mix(o.Text, v)
	if err != nil {
		return err
	}
	w.Write([]byte(text))
	w.Write([]byte{'\n'})

	return nil
}
