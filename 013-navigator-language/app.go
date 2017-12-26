package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	println(js.Global.Get("navigator").Get("language").String())
	println(js.Global.Get("navigator").Get("languages").String())
}
