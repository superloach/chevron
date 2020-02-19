package main

import (
	"syscall/js"

	"github.com/superloach/chevron"
	"github.com/superloach/chevron/errs"
)

func runF(this js.Value, _ []js.Value) interface{} {
	src_raw := src.Get("value").String()
	args := []string{}

	ch, err := chevron.Load(src_raw, args, debug.Get("checked").Bool())
	if err != nil {
		window.Call("alert", err.Error())
	}

	out.Set("value", "")

	ch.In = &inpReader{}
	ch.Out = &outWriter{}
	ch.Err = &outWriter{}

	for err == nil {
		err = ch.Step()
	}

	if !errs.Okay(err) {
		ln, lnerr := ch.Vars.Get("_#")
		if lnerr != nil {
			return nil
		}
		window.Call("alert", "error on line "+ln+": "+err.Error())
	}

	return nil
}
