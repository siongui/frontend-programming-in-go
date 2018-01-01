package main

import (
	. "github.com/siongui/godom"
)

func main() {
	Window.Call("addEventListener", "load", func(e Event) {
		s := Document.QuerySelector(".loader")
		s.ClassList().Add("invisible")
		println("All resources finished loading!")
	})
}
