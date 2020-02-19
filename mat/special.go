package mat

import (
	"time"
	"math/big"
	"math/rand"
	"strconv"
	"strings"

	"github.com/superloach/chevron/errs"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type SNumOp func(float64) (float64, error)
type STxtOp func(string) (string, error)

var SNumOps map[string]SNumOp = map[string]SNumOp{
	"p": func(n float64) (float64, error) {
		if big.NewInt(int64(n)).ProbablyPrime(0) {
			return 1, nil
		} else {
			return 0, nil
		}
	},
	"o": func(n float64) (float64, error) {
		return float64(int(n) % 2), nil
	},
	"e": func(n float64) (float64, error) {
		return float64(int(n+1) % 2), nil
	},
	"r": func(n float64) (float64, error) {
		return rand.Float64() * n, nil
	},
	"n": func(n float64) (float64, error) {
		if n == 0 {
			return 1, nil
		} else {
			return 0, nil
		}
	},
	"b": func(n float64) (float64, error) {
		if n == 0 {
			return 0, nil
		} else {
			return 1, nil
		}
	},
}

var STxtOps map[string]STxtOp = map[string]STxtOp{
	"l": func(s string) (string, error) {
		return strings.ToLower(s), nil
	},
	"u": func(s string) (string, error) {
		return strings.ToUpper(s), nil
	},
	"v": func(s string) (string, error) {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	},
	"d": func(s string) (string, error) {
		parts := strings.Split(s, "\x00")
		if len(parts) == 0 {
			return "", errs.Err("empty round")
		}

		num, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return "", err
		}

		if len(parts) == 1 {
			return strconv.FormatFloat(num, 'f', 0, 64), nil
		} else if len(parts) == 2 {
			prec, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
			return strconv.FormatFloat(num, 'f', prec, 64), nil
		} else {
			return "", errs.Err("too many round args")
		}
	},
	"c": func(s string) (string, error) {
		parts := strings.Split(s, "\x00")
		if len(parts) == 0 {
			return "", errs.Err("empty cut")
		}

		if len(parts) == 1 {
			return parts[0], nil
		} else if len(parts) == 2 {
			idx, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
			idx--
			if idx >= len(parts[0]) || idx < 0 {
				return "", errs.Err("out of bounds")
			}
			return parts[0][idx : idx+1], nil
		} else if len(parts) == 3 {
			idx1, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", err
			}
			idx1--
			if idx1 >= len(parts[0]) || idx1 < 0 {
				return "", errs.Err("out of bounds")
			}
			idx2, err := strconv.Atoi(parts[2])
			if err != nil {
				return "", err
			}
			if idx2 >= len(parts[0]) || idx2 < 0 {
				return "", errs.Err("out of bounds")
			}
			return parts[0][idx1:idx2], nil
		} else {
			return "", errs.Err("too many parts")
		}
	},
	"s": func(s string) (string, error) {
		return strconv.Itoa(len(s)), nil
	},
	"f": func(s string) (string, error) {
		parts := strings.Split(s, "\x00")
		if len(parts) != 2 {
			return "", errs.Err("wrong number of parts")
		}

		return strconv.Itoa(strings.Index(parts[0], parts[1]) + 1), nil
	},
	"a": func(s string) (string, error) {
		code, err := strconv.Atoi(s)
		if err != nil {
			return "", err
		}
		return string([]rune{rune(code)}), nil
	},
	"x": func(s string) (string, error) {
		num, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return "", err
		}
		return strconv.Itoa(int(num)), nil
	},
}
