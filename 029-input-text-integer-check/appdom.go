package main

import (
	. "github.com/siongui/godom"
	"strconv"
)

func main() {
	elmn := Document.QuerySelector("input[name='n']")
	btn := Document.QuerySelector("#check")
	info := Document.QuerySelector("#result")

	btn.AddEventListener("click", func(e Event) {
		_, err := strconv.Atoi(elmn.Value())
		if err != nil {
			info.SetInnerHTML(err.Error())
			return
		}
		info.SetInnerHTML("input is integer")
	})
}
