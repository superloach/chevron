package main

import "syscall/js"

var window js.Value = js.Global()
var document js.Value = window.Get("document")
var body js.Value = document.Get("body")

var src js.Value
var inp js.Value
var run js.Value
var out js.Value
var link js.Value
var bytes js.Value
var exmp js.Value
var debug js.Value
var printInp js.Value
var promptInp js.Value

const defaultDebug bool = false
const defaultPrintInp bool = true
const defaultPromptInp bool = true
