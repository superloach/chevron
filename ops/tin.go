package ops

import (
	"bufio"
	"io"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type TIN struct {
	Prompt string
	Var    string
}

func (t TIN) String() string {
	return "TIN `" + t.Prompt + "` `" + t.Var + "`"
}

func (t TIN) Run(v *vars.Vars, r io.Reader, w io.Writer) error {
	text, err := mix.Mix(t.Prompt, v)
	if err != nil {
		return err
	}

	w.Write([]byte(text))

	scn := bufio.NewScanner(r)
	scn.Scan()
	err = scn.Err()
	if err != nil {
		return err
	}
	value := scn.Text()

	v.Set(t.Var, value)

	return nil
}
