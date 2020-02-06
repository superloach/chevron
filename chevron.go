package chevron

import "os"
import "fmt"
import "bufio"
import "strings"
import "strconv"
import "github.com/superloach/chevron/vars"
import "github.com/superloach/chevron/ops"

type Chevron struct {
	Src string
	Args []string
	Lines []string
	Program []ops.Op
	Linenum int
	Vars *vars.Vars
	Err error
	Debug bool
}

func Load(src string, args []string, debug bool) *Chevron {
	cv := Chevron{}

	cv.Debug = debug
	cv.Err = nil

	file, err := os.Open(src)
	if err != nil {
		cv.Err = err
		return &cv
	}
	cv.Src = src

	cv.Args = args

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		cv.Err = err
		return &cv
	}

	plines := make([]string, 0)
	program := make([]ops.Op, 0)
	for _, l := range lines {
		op := ops.Find(l)
		if op != nil {
			plines = append(plines, l)
			program = append(program, op)
		}
	}
	cv.Lines = plines
	cv.Program = program

	cv.Linenum = 1

	cv.Vars = vars.MkVars()
	cv.Vars.Set("_s", src)
	cv.Vars.Set("_#", strconv.Itoa(cv.Linenum))
	cv.Vars.Set("_g", strings.Join(cv.Args, "\x00"))

	return &cv
}

func (cv *Chevron) Step() {
	if cv.Debug {
		println("line", cv.Linenum)
	}
	if cv.Linenum > len(cv.Program) {
		cv.Err = EOP
	}
	if cv.Err != nil {
		if cv.Debug {
			println("err", cv.Err.Error())
		}
		return
	}
	cv.Vars.Set("_#", strconv.Itoa(cv.Linenum))
	if cv.Debug {
		println("line", cv.Lines[cv.Linenum - 1])
	}
	op := cv.Program[cv.Linenum - 1]
	if cv.Debug {
		println("op", op.String())
	}
	cv.Err = op.Run(cv.Vars)
	cv.Linenum++
}

var EOP error = fmt.Errorf("end of program")
