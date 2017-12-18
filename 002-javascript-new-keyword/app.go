package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	d := js.Global.Get("Date").New()
	println(d)
}
