package vars

import "github.com/superloach/chevron/errs"

type Vars struct {
	values map[string]string
}

func MkVars() *Vars {
	v := Vars{}
	v.values = Default
	return &v
}

func (v *Vars) Get(n string) (string, error) {
	name := []rune(n)

	if name[0] == '^' {
		name = name[1:]
	}

	l := 1

	if name[0] == '_' || name[0] == ':' {
		l++
	}

	name = name[:l]

	val, ok := v.values[string(name)]

	if !ok {
		return "", errs.Err("^" + string(name) + " is not defined")
	}
	return val, nil
}

func (v *Vars) Set(n, value string) {
	name := []rune(n)

	if name[0] == '^' {
		name = name[1:]
	}

	l := 1

	if name[0] == '_' || name[0] == ':' {
		l++
	}

	name = name[:l]

	v.values[string(name)] = value
}

func (v *Vars) All() map[string]string {
	return v.values
}
