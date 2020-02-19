package chevron

import (
	"bufio"
	"fmt"
	"io"
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
	Err     io.Writer
	Out     io.Writer
	In      io.Reader
	Debug   bool
}

func LoadFile(src string, args []string, debug bool) (*Chevron, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	code := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		code += l + "\n"
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return Load(code, args, debug)
}

func Load(code string, args []string, debug bool) (*Chevron, error) {
	ch := Chevron{}

	ch.Args = args
	ch.Debug = debug
	ch.Lines = make([]string, 0)
	ch.Program = make([]ops.Op, 0)
	ch.Vars = vars.MkVars()

	for _, l := range strings.Split(code, "\n") {
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

	ch.Vars.Set("_#", "1")
	ch.Vars.Set("_g", strings.Join(ch.Args, "\x00"))

	return &ch, nil
}

func (ch *Chevron) DebugPrintln(args ...interface{}) {
	if ch.Debug {
		fmt.Fprint(ch.Err, "\u26A0 "+fmt.Sprintln(args...))
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

	ch.DebugPrintln("linenum", linenum)
	ch.DebugPrintln("line", ch.Lines[linenum-1])

	op := ch.Program[linenum-1]
	ch.DebugPrintln("op", op.String())

	err = op.Run(ch.Vars, ch.In, ch.Out)
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
