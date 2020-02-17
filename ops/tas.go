package ops

import (
	"io"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type TAS struct {
	Text string
	Var  string
}

func (t TAS) String() string {
	return "TAS `" + t.Text + "` `" + t.Var + "`"
}

func (t TAS) Run(v *vars.Vars, _ io.Reader, _ io.Writer) error {
	text, err := mix.Mix(t.Text, v)
	if err != nil {
		return err
	}
	v.Set(t.Var, text)
	return nil
}
