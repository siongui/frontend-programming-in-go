package main

import (
	. "github.com/siongui/godom"
	"strconv"
	"time"
)

var tooltip *Object

// indicate if the mouse cursor is in the tooltip
var isCursorInTooltip = false

// when cursor leaves the text, the delay time to close the tooltip if the
// cursor is not in the tooltip. (milisecond)
var delay = 250 * time.Millisecond

func GetPosition(elm *Object) (x, y float64) {
	x = elm.GetBoundingClientRect().Left() + Window.PageXOffset()
	y = elm.GetBoundingClientRect().Top() + Window.PageYOffset()
	return
}

// Create and append invisible tooltip to DOM tree
func SetupTooltip() {
	tooltip = Document.CreateElement("div")
	tooltip.ClassList().Add("tooltip")
	tooltip.ClassList().Add("invisible")

	tooltip.AddEventListener("mouseenter", func(e Event) {
		isCursorInTooltip = true
	})
	tooltip.AddEventListener("mouseleave", func(e Event) {
		isCursorInTooltip = false
		tooltip.ClassList().Add("invisible")
	})

	Document.QuerySelector("body").AppendChild(tooltip)
}

// event handler for mouseenter event of elements with data-descr attribute
func ShowTooltip(e Event) {
	elm := e.Target()
	tooltip.SetInnerHTML(elm.Dataset().Get("descr").String())

	x, y := GetPosition(elm)
	xpx := strconv.FormatFloat(x, 'E', -1, 64) + "px"
	ypx := strconv.FormatFloat(y+elm.Get("OffsetHeight").Float()+5, 'E', -1, 64) + "px"
	tooltip.Style().SetLeft(xpx)
	tooltip.Style().SetTop(ypx)

	tooltip.ClassList().Remove("invisible")
}

// event handler for mouseleave event of elements with data-descr attribute
func HideTooltip(e Event) {
	time.AfterFunc(delay, func() {
		if !isCursorInTooltip {
			tooltip.ClassList().Add("invisible")
		}
	})
}

func main() {
	SetupTooltip()

	// select all elements with data-descr attribute and
	// attach mouseenter and mouseleave event handler
	els := Document.QuerySelectorAll("*[data-descr]")
	for _, el := range els {
		el.AddEventListener("mouseenter", ShowTooltip)
		el.AddEventListener("mouseleave", HideTooltip)
	}
}
