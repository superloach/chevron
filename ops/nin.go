package ops

import (
	"bufio"
	"io"
	"strconv"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type NIN struct {
	Prompt string
	Var    string
}

func (n NIN) String() string {
	return "NIN `" + n.Prompt + "` `" + n.Var + "`"
}

func (n NIN) Run(v *vars.Vars, r io.Reader, w io.Writer) error {
	text, err := mix.Mix(n.Prompt, v)
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
	expr := scn.Text()

	val, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return err
	}
	v.Set(n.Var, strconv.FormatFloat(val, 'f', -1, 64))

	return nil
}
