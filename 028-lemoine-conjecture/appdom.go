package main

import (
	. "github.com/siongui/godom"
	"strconv"
)

func main() {
	elmn := Document.QuerySelector("input[name='n']")
	btn := Document.QuerySelector("#getLemoine")
	resultElm := Document.QuerySelector("#result")

	btn.AddEventListener("click", func(e Event) {
		n, err := strconv.Atoi(elmn.Value())
		if err != nil {
			resultElm.SetValue(err.Error())
			return
		}
		pqs, err := Lemoine(n)
		if err != nil {
			resultElm.SetValue(err.Error())
			return
		}

		text := ""
		for p, q := range pqs {
			text += elmn.Value() + " = " + strconv.Itoa(p) + " + ( 2 * " + strconv.Itoa(q) + " )\n"
			resultElm.SetValue(text)
		}
	})
}
