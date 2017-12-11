package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")

	f.AddEventListener("click", func(e Event) {
		f.SetTextContent("I am clicked")
	})
}
