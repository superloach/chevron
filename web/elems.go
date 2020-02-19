package main

import "syscall/js"

func br() js.Value {
	return document.Call("createElement", "br")
}

func elems() {
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
