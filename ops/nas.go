package ops

import (
	"io"

	"github.com/superloach/chevron/mat"
	"github.com/superloach/chevron/vars"
)

type NAS struct {
	Expr string
	Var  string
}

func (n NAS) String() string {
	return "NAS `" + n.Expr + "` `" + n.Var + "`"
}

func (n NAS) Run(v *vars.Vars, _ io.Reader, _ io.Writer) error {
	val, err := mat.Mat(n.Expr, v)
	if err != nil {
		return err
	}
	v.Set(n.Var, val)
	return nil
}
