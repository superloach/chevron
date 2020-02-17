package ops

import (
	"io"
	"strings"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type JMP struct {
	Rel string
	To  string
	Lh  string
	Cmp string
	Rh  string
}

func (j JMP) String() string {
	return "JMP `" + j.Rel + "` `" + j.To + "` `" + j.Lh + "` `" + j.Cmp + "` `" + j.Rh + "`"
}

type Cmp func(string, string) bool

var Cmps map[string]Cmp = map[string]Cmp{
	"=": func(l string, r string) bool {
		return l == r
	},
	"<": func(l string, r string) bool {
		return l < r
	},
	">": func(l string, r string) bool {
		return l > r
	},
	"~": func(l string, r string) bool {
		return strings.Contains(r, l)
	},
}

func AllCmps() string {
	r := ""
	for c := range Cmps {
		r += c
	}
	return r
}

func (j JMP) Run(v *vars.Vars, r io.Reader, w io.Writer) error {
	lhs, err := mix.Mix(j.Lh, v)
	if err != nil {
		return err
	}

	rhs, err := mix.Mix(j.Rh, v)
	if err != nil {
		return err
	}

	cmp, ok := Cmps[j.Cmp]
	if !ok {
		return errs.Err("unknown cmp " + j.Cmp)
	}

	if cmp(lhs, rhs) {
		return HOP{j.Rel, j.To}.Run(v, r, w)
	}

	return nil
}
