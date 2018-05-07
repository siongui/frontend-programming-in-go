package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/oskca/gopherjs-vue"
)

type Model struct {
	*js.Object        // this is needed for bidirectional data bindings
	UserInput  string `js:"userinput"`
	OldValue   string `js:"oldvalue"`
	NewValue   string `js:"newvalue"`
}

func main() {
	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// field assignment is required in this way to make data passing works
	m.UserInput = ""
	m.OldValue = ""
	m.NewValue = ""

	// create the VueJS viewModel using a struct pointer
	app := vue.New("#vueapp", m)

	option := js.Global.Get("Object").New()
	option.Set("immediate", true)
	app.Call("$watch", "userinput", func(newVal, oldVal string) {
		m.OldValue = oldVal
		m.NewValue = newVal
	}, option)
}
