package main

import (
	"github.com/gopherjs/gopherjs/js"
)

var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json"

func GetJSON(url string) {
	req := js.Global.Get("XMLHttpRequest").New()
	req.Call("addEventListener", "load", func() {
		println(req.Get("responseText").String())
	})
	req.Call("open", "GET", url, true)
	req.Call("send")
}

func main() {
	GetJSON(URL)
}
