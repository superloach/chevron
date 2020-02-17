package ops

import (
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

func (t TAS) Run(v *vars.Vars) error {
	text, err := mix.Mix(t.Text, v)
	if err != nil {
		return err
	}
	v.Set(t.Var, text)
	return nil
}
