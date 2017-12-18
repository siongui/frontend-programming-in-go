package main

import (
	. "github.com/siongui/godom"
)

func main() {
	f := Document.QuerySelector("#foo")

	f.ClassList().Add("invisible")
}
