package main

import (
	"strings"
	"syscall/js"

	"github.com/superloach/chevron"
)

func br() js.Value {
	window := js.Global()
	document := window.Get("document")
	return document.Call("createElement", "br")
}

var window js.Value = js.Global()
var document js.Value = window.Get("document")
var body js.Value = document.Get("body")

var src js.Value
var inp js.Value
var run js.Value
var out js.Value

func runF(this js.Value, _ []js.Value) interface{} {
	code := src.Get("value").String()
	args := []string{}
	debug := false

	ch, err := chevron.Load(code, args, debug)
	if err != nil {
		window.Call("alert", err.Error())
	}

	ch.In = strings.NewReader(inp.Get("value").String())
	stdout := &strings.Builder{}
	ch.Out = stdout

	for err == nil {
		err = ch.Step()
		out.Set("value", stdout.String())
	}

	return nil
}

func init() {
	src = document.Call("createElement", "textarea")

	inp = document.Call("createElement", "textarea")

	run = document.Call("createElement", "input")
	run.Set("type", "button")
	run.Set("value", "run")
	run.Call("addEventListener", "click", js.FuncOf(runF))

	out = document.Call("createElement", "textarea")
	out.Set("readOnly", "true")

	body.Call("append", "source:")
	body.Call("append", br())
	body.Call("append", src)
	body.Call("append", br())

	body.Call("append", "stdin:")
	body.Call("append", br())
	body.Call("append", inp)
	body.Call("append", br())

	body.Call("append", run)
	body.Call("append", br())

	body.Call("append", "stdout:")
	body.Call("append", br())
	body.Call("append", out)
	body.Call("append", br())
}

func main() {
	select {}
}
