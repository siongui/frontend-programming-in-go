package main

import (
	"time"
)

func main() {
	time.AfterFunc(3*time.Second, func() {
		println("3 seconds timeout")
	})
}
