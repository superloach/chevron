FROM golang:1.14
RUN go get -d -v -x github.com/superloach/chevron
WORKDIR /go/src/github.com/superloach/chevron/web
RUN GOOS=js GOARCH=wasm go build -v -x -o main.wasm .

FROM nginx:latest AS chevron
WORKDIR /usr/share/nginx/html
COPY --from=0 /go/src/github.com/superloach/chevron/web/main.wasm .
COPY --from=0 /go/src/github.com/superloach/chevron/web/index.html .
COPY --from=0 /go/src/github.com/superloach/chevron/examples .
COPY --from=0 /usr/local/go/misc/wasm/wasm_exec.js .
RUN echo "application/wasm wasm;" >> /etc/nginx/mime.types
