package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

var (
	elapsed time.Duration
)

func main() {

	fasthttp.ListenAndServe(":8081", requestHandler)
}

func fooHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q, elapsed = %s", ctx.RequestURI(), elapsed)
}
func barHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Hi there! RequestURI is %q, elapsed = %s", ctx.RequestURI(), elapsed)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	start := time.Now()
	switch string(ctx.Path()) {
	case "/foo":
		fooHandler(ctx)
	case "/bar":
		barHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
	elapsed = time.Since(start)
	log.Printf("Processing took %s", elapsed)
}
