package main

import (
	"encoding/base64"
	"net/url"
	"syscall/js"
)

func linkF(this js.Value, _ []js.Value) interface{} {
	go func() {
		src_raw := src.Get("value").String()
		inp_raw := inp.Get("value").String()

		enc := base64.URLEncoding
		src64 := enc.EncodeToString([]byte(src_raw))
		inp64 := enc.EncodeToString([]byte(inp_raw))

		raw_href := window.Get("location").Get("href").String()
		href, err := url.Parse(raw_href)
		if err != nil {
			return
		}

		query := href.Query()
		query.Set("src", src64)
		query.Set("inp", inp64)

		href.RawQuery = query.Encode()
		window.Get("location").Set("href", href.String())
	}()

	return nil
}
