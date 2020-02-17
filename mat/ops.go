package mat

import "math"

type NumOp func(float64, float64) (float64, error)

func btof(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

var NumOps map[rune]NumOp = map[rune]NumOp{
	'+': func(l float64, r float64) (float64, error) {
		return l + r, nil
	},
	'-': func(l float64, r float64) (float64, error) {
		return l - r, nil
	},
	'/': func(l float64, r float64) (float64, error) {
		return l / r, nil
	},
	'*': func(l float64, r float64) (float64, error) {
		return l * r, nil
	},
	'%': func(l float64, r float64) (float64, error) {
		return math.Mod(l, r), nil
	},
	'<': func(l float64, r float64) (float64, error) {
		return btof(l < r), nil
	},
	'=': func(l float64, r float64) (float64, error) {
		return btof(l == r), nil
	},
	'>': func(l float64, r float64) (float64, error) {
		return btof(l > r), nil
	},
	'`': func(l float64, r float64) (float64, error) {
		return math.Pow(l, r), nil
	},
}

func AllNumOps() string {
	r := make([]rune, len(NumOps))
	i := 0
	for o := range NumOps {
		r[i] = o
		i++
	}
	return string(r)
}
