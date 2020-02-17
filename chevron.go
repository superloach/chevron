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
	cv := Chevron{}

	cv.Src = src
	cv.Args = args
	cv.Debug = debug
	cv.Lines = make([]string, 0)
	cv.Program = make([]ops.Op, 0)
	cv.Vars = vars.MkVars()

	file, err := os.Open(cv.Src)
	if err != nil {
		return &cv, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		op := ops.Find(l)
		if op != nil {
			switch op.(type) {
			case ops.LBL:
				lbl := op.(ops.LBL)
				val := strconv.Itoa(len(cv.Lines) + 1)
				cv.Vars.Set(":"+lbl.Name, val)
			default:
			case ops.BAD:
				return &cv, op.(ops.BAD)
			}
			cv.Lines = append(cv.Lines, l)
			cv.Program = append(cv.Program, op)
		}
	}

	err = scanner.Err()
	if err != nil {
		return &cv, err
	}

	cv.Vars.Set("_s", src)
	cv.Vars.Set("_#", "1")
	cv.Vars.Set("_g", strings.Join(cv.Args, "\x00"))

	return &cv, nil
}

func (cv *Chevron) DebugPrint(args ...interface{}) {
	if cv.Debug {
		fmt.Println(args...)
	}
}

func (cv *Chevron) Step() error {
	lns, err := cv.Vars.Get("_#")
	if err != nil {
		return err
	}

	linenum, err := strconv.Atoi(lns)
	if err != nil {
		return err
	}

	if linenum > len(cv.Program) {
		return errs.EOF
	}

	cv.DebugPrint("linenum", linenum)
	cv.DebugPrint("line", cv.Lines[linenum-1])

	op := cv.Program[linenum-1]
	cv.DebugPrint("op", op.String())

	err = op.Run(cv.Vars)
	if err != nil {
		return err
	}

	lns = strconv.Itoa(linenum)

	nlns, err := cv.Vars.Get("_#")
	if err != nil {
		return err
	}

	if lns == nlns {
		linenum++
		cv.Vars.Set("_#", strconv.Itoa(linenum))
	}

	return err
}
