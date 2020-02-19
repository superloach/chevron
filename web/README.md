chevron web
===========
web interface for chevron.

to prepare `main.wasm`, run:
```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

if you want to be able to load examples, your web server must provide a `/examples` endpoint which returns json data similar to:
```json
[
	{"name": "foo.ch"},
	{"name": "bar.ch"},
	{"name": "baz.ch"}
]
```
(nginx has this with `autoindex on` and `autoindex_format json`)
