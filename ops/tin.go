package ops

import (
	"bufio"
	"os"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

var TINScanner = bufio.NewScanner(os.Stdin)

type TIN struct {
	Prompt string
	Var    string
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

	TINScanner.Scan()
	err = TINScanner.Err()
	if err != nil {
		return err
	}
	value := Clean(TINScanner.Text())

	v.Set(t.Var, value)

	return nil
}
