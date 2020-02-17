package main

import (
	"io"
	"syscall/js"

	"github.com/superloach/chevron"
	"github.com/superloach/chevron/errs"
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

type textareaReader struct {
	textarea js.Value
	index int
}

func (r *textareaReader) Read(p []byte) (int, error) {
	i := 0
	v := r.textarea.Get("value").String()
	for i < len(p) && r.index + i < len(v) {
		println(r.index + i, len(v))
		b := v[r.index + i]
		if b == '\n' {
			r.index += i + 1
			return i, io.EOF
		}
		p[i] = b
		i++
	}
	r.index += i + 1
	return i, io.EOF
}

type textareaWriter struct {
	textarea js.Value
}

func (w *textareaWriter) Write(p []byte) (int, error) {
	v := w.textarea.Get("value").String()
	v += string(p)
	w.textarea.Set("value", v)
	return len(p), nil
}

func runF(this js.Value, _ []js.Value) interface{} {
	code := src.Get("value").String()
	args := []string{}
	debug := false

	ch, err := chevron.Load(code, args, debug)
	if err != nil {
		window.Call("alert", err.Error())
	}

	stdin := &textareaReader{inp, 0}
	ch.In = stdin

	stdout := &textareaWriter{out}
	ch.Out = stdout

	for err == nil {
		err = ch.Step()
	}

	if !errs.Okay(err) {
		ln, lnerr := ch.Vars.Get("_#")
		if lnerr != nil {
			panic(lnerr)
		}
		window.Call("alert", "error on line " + ln + ": " + err.Error())
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
