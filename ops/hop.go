package ops

import (
	"io"
	"strconv"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/mat"
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

type HOP struct {
	Rel string
	To  string
}

func (h HOP) String() string {
	return "HOP `" + h.Rel + "` `" + h.To + "`"
}

func (h HOP) Run(v *vars.Vars, _ io.Reader, _ io.Writer) error {
	val, err := mix.Mix(h.To, v)
	if err != nil {
		return err
	}

	lbl, _ := v.Get(":" + val)
	if lbl != "" {
		val = lbl
	}
	val, err = mat.Mat(val, v)
	if err != nil {
		return err
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
	value := strconv.Itoa(int(num))
	v.Set("_#", value)
	return nil
}
