package main

import (
	. "github.com/siongui/godom"
	"regexp"
)

var textUrl = map[string]string{
	"世尊譯詞的探討": "http://agama.buddhason.org/book/bb/bb21.htm",
}

var text = regexp.MustCompile(`〈.+〉`)

func TextToLink(elm *Object) {
	h := text.ReplaceAllStringFunc(elm.InnerHTML(), func(match string) string {
		key := match[3 : len(match)-3]
		url, ok := textUrl[key]
		if ok {
			return `〈<a href="` + url +
				`" target="_blank">` + key + `</a>〉`
		}
		return match
	})
	elm.SetInnerHTML(h)
}

func main() {
	b := Document.QuerySelector(".line-block")
	TextToLink(b)
}
