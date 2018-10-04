package main

import (
	. "github.com/siongui/godom/wasm"
)

func main() {
	testdiv := Document.QuerySelector("#testdiv")
	testdiv.Set("innerHTML", "hi")
}
