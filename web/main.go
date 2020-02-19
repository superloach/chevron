package main

import (
	"io"
	"net/url"
	"syscall/js"
	"encoding/base64"

	"github.com/superloach/chevron"
	"github.com/superloach/chevron/errs"
)

var window js.Value = js.Global()
var document js.Value = window.Get("document")
var body js.Value = document.Get("body")

var src js.Value
var inp js.Value
var run js.Value
var link js.Value
var out js.Value

func br() js.Value {
	return document.Call("createElement", "br")
}

type inpReader struct {
	index int
}

func (r *inpReader) Read(p []byte) (int, error) {
	i := 0
	v := inp.Get("value").String()
	if r.index + i >= len(v) {
		prompt := window.Call("prompt", "")
		if prompt.Truthy() {
			v = prompt.String()
		}
	}
	for i < len(p) && r.index + i < len(v) {
		println(r.index + i, len(v))
		b := v[r.index + i]
		out.Set("value", out.Get("value").String() + string(b))
		if b == '\n' {
			break
		}
		p[i] = b
		i++
	}
	if r.index + i >= len(v) {
		out.Set("value", out.Get("value").String() + "\n")
	}
	r.index += i + 1
	return i, io.EOF
}

type outWriter struct {}

func (w *outWriter) Write(p []byte) (int, error) {
	v := out.Get("value").String()
	v += string(p)
	out.Set("value", v)
	return len(p), nil
}

func runF(this js.Value, _ []js.Value) interface{} {
	src_raw := src.Get("value").String()
	args := []string{}
	debug := false

	ch, err := chevron.Load(src_raw, args, debug)
	if err != nil {
		window.Call("alert", err.Error())
	}

	out.Set("value", "")

	ch.Out = &outWriter{}
	ch.In = &inpReader{}

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

func linkF(this js.Value, _ []js.Value) interface{} {
	src_raw := src.Get("value").String()
	inp_raw := inp.Get("value").String()

	enc := base64.URLEncoding
	src64 := enc.EncodeToString([]byte(src_raw))
	inp64 := enc.EncodeToString([]byte(inp_raw))

	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	if err != nil {
		panic(err)
	}

	query := href.Query()
	query.Set("src", src64)
	query.Set("inp", inp64)
	href.RawQuery = query.Encode()
	window.Get("location").Set("href", href.String())

	return nil
}

func init() {
	src = document.Call("createElement", "textarea")

	inp = document.Call("createElement", "textarea")

	run = document.Call("createElement", "input")
	run.Set("type", "button")
	run.Set("value", "run")
	run.Call("addEventListener", "click", js.FuncOf(runF))

	link = document.Call("createElement", "input")
	link.Set("type", "button")
	link.Set("value", "link")
	link.Call("addEventListener", "click", js.FuncOf(linkF))

	out = document.Call("createElement", "textarea")
	out.Set("readOnly", "true")

	body.Call("append", "src:")
	body.Call("append", br())
	body.Call("append", src)
	body.Call("append", br())

	body.Call("append", "inp:")
	body.Call("append", br())
	body.Call("append", inp)
	body.Call("append", br())

	body.Call("append", run)
	body.Call("append", link)
	body.Call("append", br())

	body.Call("append", "out:")
	body.Call("append", br())
	body.Call("append", out)
	body.Call("append", br())

	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	if err != nil {
		panic(err)
	}

	query := href.Query()
	src64 := query.Get("src")
	inp64 := query.Get("inp")

	enc := base64.URLEncoding
	src_raw, _ := enc.DecodeString(src64)
	src.Set("value", string(src_raw))
	inp_raw, _ := enc.DecodeString(inp64)
	inp.Set("value", string(inp_raw))
}

func main() {
	select {}
}
