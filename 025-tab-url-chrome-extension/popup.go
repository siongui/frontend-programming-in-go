package main

import (
	"github.com/fabioberger/chrome"
	. "github.com/siongui/godom"
)

func main() {
	c := chrome.NewChrome()
	queryInfo := chrome.Object{
		"active":        true,
		"currentWindow": true,
	}
	c.Tabs.Query(queryInfo, func(tabs []chrome.Tab) {
		tab := tabs[0]
		Document.Write(tab.Url)
	})
}
