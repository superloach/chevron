package main

import (
	"encoding/base64"
	"io"
	"net/url"
	"strconv"
	"strings"
	"syscall/js"

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
var bytes js.Value
var out js.Value
var debug js.Value
var printInp js.Value
var promptInp js.Value

const defaultDebug bool = false
const defaultPrintInp bool = true
const defaultPromptInp bool = true

func br() js.Value {
	return document.Call("createElement", "br")
}

type inpReader struct {
	index int
}

func (r *inpReader) Read(p []byte) (int, error) {
	i := 0
	v := inp.Get("value").String()
	if r.index+i >= len(v) && promptInp.Get("checked").Bool() {
		ls := strings.Split(out.Get("value").String(), "\n")
		last := ls[len(ls)-1]
		prompt := window.Call("prompt", last)
		if prompt.Truthy() {
			v = prompt.String()
		}
	}
	for i < len(p) && r.index+i < len(v) {
		b := v[r.index+i]
		if printInp.Get("checked").Bool() {
			out.Set("value", out.Get("value").String()+string(b))
		}
		if b == '\n' {
			break
		}
		p[i] = b
		i++
	}
	if r.index+i >= len(v) {
		out.Set("value", out.Get("value").String()+"\n")
	}
	r.index += i + 1
	return i, io.EOF
}

type outWriter struct{}

func (w *outWriter) Write(p []byte) (int, error) {
	v := out.Get("value").String()
	v += string(p)
	out.Set("value", v)
	return len(p), nil
}

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
			panic(lnerr)
		}
		window.Call("alert", "error on line "+ln+": "+err.Error())
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

func bytesF(this js.Value, _ []js.Value) interface{} {
	src_raw := src.Get("value").String()

	src_bytes := strconv.Itoa(len([]byte(src_raw)))
	src_runes := strconv.Itoa(len([]rune(src_raw)))
	src_lines := strconv.Itoa(len(strings.Split(src_raw, "\n")))

	window.Call("alert", src_bytes+" bytes\n"+src_runes+" runes\n"+src_lines+" lines")

	return nil
}

func mkElements() {
	src = document.Call("createElement", "textarea")
	body.Call("append", "src:")
	body.Call("append", br())
	body.Call("append", src)
	body.Call("append", br())

	inp = document.Call("createElement", "textarea")
	body.Call("append", "inp:")
	body.Call("append", br())
	body.Call("append", inp)
	body.Call("append", br())

	run = document.Call("createElement", "input")
	run.Set("type", "button")
	run.Set("value", "run")
	run.Call("addEventListener", "click", js.FuncOf(runF))
	body.Call("append", run)

	link = document.Call("createElement", "input")
	link.Set("type", "button")
	link.Set("value", "link")
	link.Call("addEventListener", "click", js.FuncOf(linkF))
	body.Call("append", link)

	bytes = document.Call("createElement", "input")
	bytes.Set("type", "button")
	bytes.Set("value", "bytes")
	bytes.Call("addEventListener", "click", js.FuncOf(bytesF))
	body.Call("append", bytes)
	body.Call("append", br())

	out = document.Call("createElement", "textarea")
	out.Set("readOnly", true)
	body.Call("append", "out:")
	body.Call("append", br())
	body.Call("append", out)
	body.Call("append", br())

	debug = document.Call("createElement", "input")
	debug.Set("type", "checkbox")
	debug.Set("checked", defaultDebug)
	body.Call("append", debug)
	body.Call("append", "debug")
	body.Call("append", br())

	printInp = document.Call("createElement", "input")
	printInp.Set("type", "checkbox")
	printInp.Set("checked", defaultPrintInp)
	body.Call("append", printInp)
	body.Call("append", "print inp")
	body.Call("append", br())

	promptInp = document.Call("createElement", "input")
	promptInp.Set("type", "checkbox")
	promptInp.Set("checked", defaultPromptInp)
	body.Call("append", promptInp)
	body.Call("append", "prompt inp")
}

func parseQuery() {
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
	inp_raw, _ := enc.DecodeString(inp64)

	src.Set("value", string(src_raw))
	inp.Set("value", string(inp_raw))
}

func loaded() {
	loading := document.Call("getElementById", "loading")
	if loading.Truthy() {
		loading.Call("remove")
	}
}

func main() {
	mkElements()
	parseQuery()

	loaded()

	select {}
}
