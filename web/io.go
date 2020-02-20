package main

import (
	"io"
	"strings"
)

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