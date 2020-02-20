package main

import "encoding/json"
import "net/url"
import "syscall/js"
import "fmt"

func examples() []string {
	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	if err != nil {
		panic(err)
	}

	href.Path = "/examples/"
	hrefs := href.String()

	exs := make([]string, 0)
	done := make(chan struct{})

	window.Call("fetch", hrefs).Call("then",
		js.FuncOf(func (this js.Value, args []js.Value) interface{} {
			args[0].Call("text").Call("then",
				js.FuncOf(func (this js.Value, args []js.Value) interface{} {
					data := []byte(args[0].String())
					exs_raw := make([]map[string]interface{}, 0)
					json.Unmarshal(data, &exs_raw)
					for _, exr := range exs_raw {
						iname, ok := exr["name"]
						if !ok {
							continue
						}
						name, ok := iname.(string)
						if !ok {
							continue
						}
						exs = append(exs, name)
					}
					done <- struct{}{}
					return nil
				}),
			)
			return nil
		}),
	)

	<-done
	return exs
}

func exmpF(this js.Value, _ []js.Value) interface{} {
	name := exmp.Get("value").String()

	raw_href := window.Get("location").Get("href").String()
	href, err := url.Parse(raw_href)
	if err != nil {
		panic(err)
	}

	href.Path = "/examples/" + name
	href.RawQuery = ""
	hrefs := href.String()

	window.Call("fetch", hrefs).Call("then",
		js.FuncOf(func (this js.Value, args []js.Value) interface{} {
			if len(args) == 0 {
				return nil
			}
			if !args[0].Get("ok").Bool() {
				return nil
			}
			args[0].Call("text").Call("then",
				js.FuncOf(func (this js.Value, args []js.Value) interface{} {
					src.Set("value", args[0].String())
					return nil
				}),
			)
			return nil
		}),
	)

	return nil
}
