package main

import (
	. "github.com/siongui/godom"
	"strconv"
)

var tooltip *Object

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
	tooltip.ClassList().Add("invisible")
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
