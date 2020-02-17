package ops

import (
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type OUT struct {
	Text string
}

func (o OUT) String() string {
	return "OUT `" + o.Text + "`"
}

func (o OUT) Run(v *vars.Vars) error {
	text, err := mix.Mix(o.Text, v)
	if err != nil {
		return err
	}
	println(text)
	return nil
}
