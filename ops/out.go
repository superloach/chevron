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

	if text[len(text) - 1] == '\x00' {
		text = text[:len(text) - 1]
	} else {
		text += "\n"
	}

	w.Write([]byte(text))

	return nil
}
