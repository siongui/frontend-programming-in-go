package main

import (
	. "github.com/siongui/godom"
)

func SetupNavbarBurgers() {
	nbs := Document.QuerySelectorAll(".navbar-burger")

	for _, nb := range nbs {
		nb.AddEventListener("click", func(e Event) {
			targetId := nb.Dataset().Get("target").String()
			target := Document.GetElementById(targetId)

			nb.ClassList().Toggle("is-active")
			target.ClassList().Toggle("is-active")
		})
	}
}

func main() {
	Document.AddEventListener("DOMContentLoaded", func(e Event) {

		// Dropdowns in navbar
		dds := Document.QuerySelectorAll(".navbar-item.has-dropdown:not(.is-hoverable)")

		closeDropdowns := func() {
			for _, dd := range dds {
				dd.ClassList().Remove("is-active")
			}
		}

		if len(dds) > 0 {
			for _, dd := range dds {
				dd.AddEventListener("click", func(e Event) {
					e.StopPropagation()
					dd.ClassList().Toggle("is-active")
				})
			}

			Document.AddEventListener("click", func(e Event) {
				closeDropdowns()
			})
		}

		// Close dropdowns if ESC pressed
		Document.AddEventListener("keydown", func(e Event) {
			if e.KeyCode() == 27 {
				closeDropdowns()
			}
		})

		// Toggles
		SetupNavbarBurgers()
	})
}
