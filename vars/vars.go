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

func (v *Vars) Get(name string) (string, error) {
	if name[0] == '^' {
		name = name[1:]
	}

	val, ok := v.values[name]

	if !ok {
		return "", errs.Err("^" + name + " is not defined")
	}
	return val, nil
}

func (v *Vars) Set(name, value string) {
	if name[0] == '^' {
		name = name[1:]
	}

	v.values[name] = value
}

func (v *Vars) All() map[string]string {
	return v.values
}
