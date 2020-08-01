package main

import (
	"syscall/js"
	"time"

	"github.com/superloach/chevron"
	"github.com/superloach/chevron/errs"
)

var ch *chevron.Chevron
var err error

func runF() {
	src_raw := src.Get("value").String()
	args := []string{}

	ch, err = chevron.Load(src_raw, args, debug.Get("checked").Bool())
	if err != nil {
		window.Call("alert", err.Error())
		panic(err)
	}

	out.Set("value", "")

	runStop.Set("value", "stop")

	ch.In = &inpReader{}
	ch.Out = &outWriter{}
	ch.Err = &outWriter{}

	dl := delay.Get("checked").Bool()
	for err == nil {
		err = ch.Step()
		if dl {
			time.Sleep(time.Millisecond)
		}
	}

	runStop.Set("value", "run")

	if !errs.Okay(err) {
		ln, lnerr := ch.Vars.Get("_#")
		if lnerr != nil {
			panic(err)
		}
		window.Call("alert", "error on line "+ln+": "+err.Error())
	}
}

func runStopF(this js.Value, _ []js.Value) interface{} {
	go func() {
		if ch != nil && err == nil {
			err = errs.EOF
		} else {
			go runF()
		}
	}()

	return nil
}
