package ops

import "github.com/superloach/chevron/vars"

type LBL struct {
	Name string
}

func (l LBL) String() string {
	return "LBL `" + l.Name + "`"
}

func (l LBL) Run(v *vars.Vars) error {
	return nil
}
