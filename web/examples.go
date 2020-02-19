package main

import "encoding/json"
import "net/http"
import "net/url"
import "fmt"

func examples() {
	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	fmt.Println(href)
	if err != nil {
		return
	}

	href.Path = "/examples"
	resp, err := http.Get(href.String())
	fmt.Println(resp)
	if err != nil {
		return
	}

	exs := make([]map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&exs)
	fmt.Println(exs)
	if err != nil {
		return
	}
}
