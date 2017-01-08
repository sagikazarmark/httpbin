package main

//go:generate go-bindata -pkg endpoints -o endpoints/bindata.go data/...

import (
	"log"
	"net/http"

	"github.com/sagikazarmark/bingo"
	"github.com/sagikazarmark/httpbin/endpoints"
	flag "github.com/spf13/pflag"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Custom port")
}

func main() {
	bin, _ := bingo.New("httpbin", "HTTP Request & Response service", "Testing an HTTP Library can become difficult sometimes. The original (hosted) version of HTTP Bin is awesome for playing around, but not really ideal for being integrated into a CI pipeline. This clone provides a single binary which can be downloaded on the fly or shipped with the test suite. Since this is a drop-in replacement, it works with your existing code.")

	endpoints.Info(bin)
	endpoints.Verb(bin)
	endpoints.Encoding(bin)
	endpoints.Response(bin)
	endpoints.Redirect(bin)

	router := bingo.NewHandler(bin)

	log.Println("Starting server on *:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
