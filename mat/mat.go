package mat

import (
	"strconv"
	"strings"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

func Mat(expr string, v *vars.Vars) (string, error) {
	opi := strings.Index(expr, "~")
	if opi == -1 {
		opi = strings.IndexAny(expr, AllNumOps())
		if opi == -1 {
			num, err := strconv.ParseFloat(expr, 64)
			if err != nil {
				return "0", err
			}
			return strconv.FormatFloat(num, 'f', -1, 64), nil
		}
	}

	opn := rune(expr[opi])

	if opn == '~' {
		so := expr[opi+1:]
		sno, ok := SNumOps[so]
		if !ok {
			sto, ok := STxtOps[so]
			if !ok {
				return "0", errs.Err("unknown special op " + so)
			}
			txt1 := strings.ReplaceAll(expr[:opi], ",", "\x00")
			txt2, err := mix.Mix(txt1, v)
			if err != nil {
				return "0", err
			}
			txt3, err := sto(txt2)
			return txt3, err
		}
		n, err := strconv.ParseFloat(expr[:opi], 64)
		if err != nil {
			return "0", err
		}
		num, err := sno(n)
		if err != nil {
			return "0", err
		}
		return strconv.FormatFloat(num, 'f', -1, 64), nil
	}

	op, ok := NumOps[opn]

	if !ok {
		return "0", errs.Err("unknown mat op " + string(opn))
	}

	lhm, err := mix.Mix(expr[:opi], v)
	if err != nil {
		return "0", err
	}
	lhs := strings.Trim(lhm, " ")
	lh, err := strconv.ParseFloat(lhs, 64)
	if err != nil {
		return "0", err
	}

	rhm, err := mix.Mix(expr[opi+1:], v)
	if err != nil {
		return "0", err
	}
	rhs := strings.Trim(rhm, " ")
	rh, err := strconv.ParseFloat(rhs, 64)
	if err != nil {
		return "0", err
	}

	num, err := op(lh, rh)
	if err != nil {
		return "0", err
	}

	return strconv.FormatFloat(num, 'f', -1, 64), nil
}
