package ops

import "github.com/superloach/chevron/vars"
import "github.com/superloach/chevron/mix"

type TIN struct {
	Prompt string
	Var string
}

func (t TIN) String() string {
	return "TIN `" + t.Prompt + "` `" + t.Var + "`"
}

func (t TIN) Run(v *vars.Vars) error {
	text, err := mix.Mix(t.Prompt, v)
	if err != nil {
		return err
	}
	print(text)
	println("STUB INPUT")
	v.Set(t.Var, "STUB")
	return nil
}
