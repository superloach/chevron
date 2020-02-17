package ops

import (
	"bufio"
	"os"
	"strconv"

	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

var NINScanner = bufio.NewScanner(os.Stdin)

type NIN struct {
	Prompt string
	Var    string
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

	NINScanner.Scan()
	err = NINScanner.Err()
	if err != nil {
		return err
	}
	expr := NINScanner.Text()

	val, err := strconv.ParseFloat(expr, 64)
	if err != nil {
		return err
	}
	v.Set(n.Var, strconv.FormatFloat(val, 'f', -1, 64))

	return nil
}
