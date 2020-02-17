package chevron

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/superloach/chevron/errs"
	"github.com/superloach/chevron/ops"
	"github.com/superloach/chevron/vars"
)

type Chevron struct {
	Src     string
	Args    []string
	Lines   []string
	Program []ops.Op
	Vars    *vars.Vars
	Debug   bool
}

func Load(src string, args []string, debug bool) (*Chevron, error) {
	ch := Chevron{}

	ch.Src = src
	ch.Args = args
	ch.Debug = debug
	ch.Lines = make([]string, 0)
	ch.Program = make([]ops.Op, 0)
	ch.Vars = vars.MkVars()

	file, err := os.Open(ch.Src)
	if err != nil {
		return &ch, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		op := ops.Find(l)
		if op != nil {
			switch op.(type) {
			case ops.LBL:
				lbl := op.(ops.LBL)
				val := strconv.Itoa(len(ch.Lines) + 1)
				ch.Vars.Set(":"+lbl.Name, val)
			default:
			case ops.BAD:
				return &ch, op.(ops.BAD)
			}
			ch.Lines = append(ch.Lines, l)
			ch.Program = append(ch.Program, op)
		}
	}

	err = scanner.Err()
	if err != nil {
		return &ch, err
	}

	ch.Vars.Set("_s", src)
	ch.Vars.Set("_#", "1")
	ch.Vars.Set("_g", strings.Join(ch.Args, "\x00"))

	return &ch, nil
}

func (ch *Chevron) DebugPrint(args ...interface{}) {
	if ch.Debug {
		fmt.Println(args...)
	}
}

func (ch *Chevron) Step() error {
	lns, err := ch.Vars.Get("_#")
	if err != nil {
		return err
	}

	linenum, err := strconv.Atoi(lns)
	if err != nil {
		return err
	}

	if linenum > len(ch.Program) {
		return errs.EOF
	}

	ch.DebugPrint("linenum", linenum)
	ch.DebugPrint("line", ch.Lines[linenum-1])

	op := ch.Program[linenum-1]
	ch.DebugPrint("op", op.String())

	err = op.Run(ch.Vars)
	if err != nil {
		return err
	}

	lns = strconv.Itoa(linenum)

	nlns, err := ch.Vars.Get("_#")
	if err != nil {
		return err
	}

	if lns == nlns {
		linenum++
		ch.Vars.Set("_#", strconv.Itoa(linenum))
	}

	return err
}
