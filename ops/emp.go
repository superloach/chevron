package ops

import "github.com/superloach/chevron/vars"

type EMP struct{}

func (e EMP) String() string {
	return "EMP"
}

func (e EMP) Run(v *vars.Vars) error {
	return nil
}
