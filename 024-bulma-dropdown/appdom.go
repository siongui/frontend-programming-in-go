package main

import (
	. "github.com/siongui/godom"
)

func main() {
	Document.AddEventListener("DOMContentLoaded", func(e Event) {

		// Dropdowns
		dds := Document.QuerySelectorAll(".dropdown:not(.is-hoverable)")

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
	})
}
