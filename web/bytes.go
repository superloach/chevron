package main

import (
	"strconv"
	"strings"
	"syscall/js"
)

func bytesF(this js.Value, _ []js.Value) interface{} {
	go func() {
		src_raw := src.Get("value").String()

		src_bytes := strconv.Itoa(len([]byte(src_raw)))
		src_runes := strconv.Itoa(len([]rune(src_raw)))
		src_lines := strconv.Itoa(len(strings.Split(src_raw, "\n")))

		window.Call("alert", src_bytes+" bytes\n"+src_runes+" runes\n"+src_lines+" lines")
	}()

	return nil
}
