package ops

import "github.com/superloach/chevron/vars"

type Op interface {
	Run(*vars.Vars) error
	String() string
}
