package ops

import (
	"strconv"

	"github.com/superloach/chevron/mat"
	"github.com/superloach/chevron/vars"
)

type SKP struct {
	Rel string
	To string
	If string
}

func (s SKP) String() string {
	return "SKP `" + s.Rel + "` `" + s.To + "` `" + s.If + "`"
}

func (s SKP) Run(v *vars.Vars) error {
	ifs, err := mat.Mat(s.If, v)
	if err != nil {
		return err
	}

	ifn, err := strconv.ParseFloat(ifs, 64)
	if err != nil {
		return err
	}

	if ifn != 0 {
		return HOP{s.Rel, s.To}.Run(v)
	}

	return nil
}
