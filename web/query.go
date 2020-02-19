package main

import (
	"encoding/base64"
	"net/url"
)

func query() {
	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	if err != nil {
		return
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
