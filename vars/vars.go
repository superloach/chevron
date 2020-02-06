package vars

type Vars struct {
	values map[string]string
}

func MkVars() *Vars {
	v := Vars{}
	v.values = Default
	return &v
}

func (v *Vars) Get(name string) string {
	return v.values[name]
}

func (v *Vars) Set(name, value string) {
	v.values[name] = value
}
