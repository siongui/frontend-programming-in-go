package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object           // this is needed for bidirectional data bindings
	PaliWord    string   `js:"paliWord"`
	Row1letters []string `js:"row1letters"`
	Row2letters []string `js:"row2letters"`
	Row3letters []string `js:"row3letters"`
	Row4letters []string `js:"row4letters"`
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.PaliWord = ""
	m.Row1letters = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"}
	m.Row2letters = []string{"a", "s", "d", "f", "g", "h", "j", "k", "l"}
	m.Row3letters = []string{"z", "x", "c", "v", "b", "n", "m"}
	m.Row4letters = []string{"ā", "ḍ", "ī", "ḷ", "ṁ", "ṃ", "ñ", "ṇ", "ṭ", "ū", "ŋ", "ṅ"}

	// create the VueJS viewModel using a struct pointer
	vue.New("#container", m)
}
