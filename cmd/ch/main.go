package main

import (
	"flag"
	"os"

	"github.com/superloach/chevron"
	"github.com/superloach/chevron/errs"
)

func main() {
	src := flag.String("src", "", "file to run")
	debug := flag.Bool("debug", false, "increase verbosity")
	flag.Parse()
	args := flag.Args()

	switch {
	case *src == "":
		if len(args) > 0 {
			*src = args[0]
			args = args[1:]
		} else {
			flag.Usage()
			os.Exit(0)
		}
	default:
	}

	ch, err := chevron.Load(*src, args, *debug)
	if err != nil {
		panic(err)
	}

	for err == nil {
		err = ch.Step()
	}

	if !errs.Okay(err) {
		ln, lnerr := ch.Vars.Get("_#")
		if lnerr != nil {
			panic(lnerr)
		}
		panic(errs.At{ln, err})
	}
}
