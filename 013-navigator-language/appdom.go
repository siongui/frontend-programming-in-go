package main

import (
	. "github.com/siongui/godom"
)

func main() {
	println(Window.Navigator().Language())
	println(Window.Navigator().Languages())
}
