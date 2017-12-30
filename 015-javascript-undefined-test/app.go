package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	if js.Global.Get("localStorage") == js.Undefined {
		println("Your browser does not support localStorage API")
	} else {
		println("Your browser supports localStorage API")
	}
}
