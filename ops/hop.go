package ops

import (
	"strconv"

	"github.com/superloach/chevron/mat"
	"github.com/superloach/chevron/vars"
	"github.com/superloach/chevron/errs"
)

type HOP struct {
	Rel string
	To string
}

func (h HOP) String() string {
	return "HOP `" + h.Rel + "` `" + h.To + "`"
}

func (h HOP) Run(v *vars.Vars) error {
	val := "0"
	val, err := mat.Mat(h.To, v)
	if err != nil {
		lbl, _ := v.Get(":" + h.To)
		if lbl == "" {
			return err
		}
		val, err = mat.Mat(lbl, v)
		if err != nil {
			return err
		}
	}
	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return err
	}
	switch h.Rel {
	case "":
		break
	case "+":
		cs, err := v.Get("_#")
		if err != nil {
			return err
		}
		cur, err := strconv.ParseFloat(cs, 64)
		if err != nil {
			return err
		}
		num = cur + num
	case "-":
		cs, err := v.Get("_#")
		if err != nil {
			return err
		}
		cur, err := strconv.ParseFloat(cs, 64)
		if err != nil {
			return err
		}
		num = cur - num
	default:
		return errs.Err("unknown relative operator")
	}
	val = strconv.Itoa(int(num))
	v.Set("_#", val)
	return nil
}
