package ops

import "github.com/superloach/chevron/vars"

type BAD struct {
	Line string
	Reason string
}

func (b BAD) String() string {
	return "BAD `" + b.Line + "` `" + b.Reason + "`"
}

func (b BAD) Run(v *vars.Vars) error {
	return nil
}

func (b BAD) Error() string {
	return "bad line: " + b.Line + " (" + b.Reason + ")"
}
