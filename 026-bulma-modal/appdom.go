package main

import (
	. "github.com/siongui/godom"
)

func main() {
	Document.AddEventListener("DOMContentLoaded", func(e Event) {

		// Modals

		rootEl := Document.Get("documentElement")
		modals := Document.QuerySelectorAll(".modal")
		modalButtons := Document.QuerySelectorAll(".modal-button")
		modalCloses := Document.QuerySelectorAll(".modal-background, .modal-close, .modal-card-head .delete, .modal-card-foot .button")

		closeModals := func() {
			rootEl.Get("classList").Call("remove", "is-clipped")
			for _, modal := range modals {
				modal.ClassList().Remove("is-active")
			}
		}

		if len(modalButtons) > 0 {
			for _, modalButton := range modalButtons {
				modalButton.AddEventListener("click", func(e Event) {
					targetId := modalButton.Dataset().Get("target").String()
					target := Document.GetElementById(targetId)
					rootEl.Get("classList").Call("add", "is-clipped")
					target.ClassList().Add("is-active")
				})
			}
		}

		if len(modalCloses) > 0 {
			for _, modalClose := range modalCloses {
				modalClose.AddEventListener("click", func(e Event) {
					closeModals()
				})
			}
		}

		// Close modals if ESC pressed
		Document.AddEventListener("keydown", func(e Event) {
			if e.KeyCode() == 27 {
				closeModals()
			}
		})
	})
}
