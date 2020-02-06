package ops

import "github.com/superloach/chevron/vars"
import "github.com/superloach/chevron/mat"

type NAS struct {
	Expr string
	Var string
}

func (n NAS) String() string {
	return "NAS `" + n.Var + "` `" + n.Expr + "`"
}

func (n NAS) Run(v *vars.Vars) error {
	val, err := mat.Mat(n.Expr, v)
	if err != nil {
		return err
	}
	v.Set(n.Var, val)
	return nil
}
