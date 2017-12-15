package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")

	// get value
	println(f.Dataset().Get("demoValue").String())

	// set value
	f.Dataset().Set("demoValue", "world hello")
	println(f.Dataset().Get("demoValue").String())
}
