package ops

import (
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

func (d DIE) Run(v *vars.Vars) error {
	text, err := mix.Mix(d.Text, v)
	if err != nil {
		return err
	}
	println(text)
	return errs.DIE
}
