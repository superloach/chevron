package ops

import "github.com/superloach/chevron/vars"
import "github.com/superloach/chevron/mix"

type NIN struct {
	Prompt string
	Var string
}

func (n NIN) String() string {
	return "NIN `" + n.Prompt + "` `" + n.Var + "`"
}

func (n NIN) Run(v *vars.Vars) error {
	text, err := mix.Mix(n.Prompt, v)
	if err != nil {
		return err
	}
	print(text)
	println("STUB INPUT")
	v.Set(n.Var, "STUB")
	return nil
}
