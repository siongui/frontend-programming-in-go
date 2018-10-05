package main

import (
	. "github.com/siongui/godom/wasm"
)

func main() {
	testdivs := Document.QuerySelectorAll("#testdivs > div")
	for _, testdiv := range testdivs {
		testdiv.Set("innerHTML", "hi")
	}
}
