package main

import (
	"bytes"
	"net/http"
)

var URL = "https://siongui.github.io/xemaauj9k5qn34x88m4h/sacca.json"

func GetJSON(url string) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	println(buf.String())
}

func main() {
	GetJSON(URL)
}
