package mat

import (
	"strconv"
	"strings"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/mix"
	"github.com/superloach/chevron/vars"
)

func Mat(expr string, v *vars.Vars) (string, error) {
	expr, err := mix.Mix(expr, v)
	if err != nil {
		return "0", err
	}

	opi := strings.IndexAny(expr, AllNumOps()+"~")
	if opi == -1 {
		num, err := strconv.ParseFloat(expr, 64)
		if err != nil {
			return "0", err
		}
		return strconv.FormatFloat(num, 'f', -1, 64), nil
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
			txt, err := sto(expr[:opi])
			return txt, err
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

	lhs := strings.Trim(expr[:opi], " ")
	lh, err := strconv.ParseFloat(lhs, 64)
	if err != nil {
		return "0", err
	}

	rhs := strings.Trim(expr[opi+1:], " ")
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
