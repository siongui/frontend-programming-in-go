package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object        // this is needed for bidirectional data bindings
	Keypressed string `js:"keypressed"`
}

// this would be recognized as Show in html
func (m *Model) ShowKeyCode(event *js.Object) {
	m.Keypressed = event.Get("keyCode").String()
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.Keypressed = ""
	// create the VueJS viewModel using a struct pointer
	vue.New("#vueapp", m)
}
