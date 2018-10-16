package main

import (
	"fmt"
	"syscall/js"

	. "github.com/siongui/godom/wasm"
)

var signal = make(chan int)

func keepAlive() {
	for {
		<-signal
	}
}

func main() {
	foo := Document.GetElementById("foo")
	count := 0

	cb := js.NewCallback(func(args []js.Value) {
		count++
		foo.Set("textContent", fmt.Sprintf("I am clicked %d time", count))
	})
	foo.Call("addEventListener", "click", cb)

	keepAlive()
}
