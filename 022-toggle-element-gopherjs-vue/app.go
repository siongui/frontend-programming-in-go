package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object      // this is needed for bidirectional data bindings
	IsShow     bool `js:"isShow"`
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.IsShow = false
	// create the VueJS viewModel using a struct pointer
	vue.New("#app", m)
}
