package wg

import (
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

func GetElementById(id string) js.Value {
	return Doc.Call("getElementById", id)
}

func CreateElementFromHTML(html string) js.Value {
	div := Doc.Call("createElement", "div")
	div.Set("innerHTML", strings.TrimSpace(html))
	return div.Get("firstChild")
}

func AppendBody(htmlElement js.Value) {
	Body.Call("appendChild", htmlElement)
}

func Display(htmlElement js.Value, val string) {
	style := htmlElement.Get("style")
	style.Set("display", val)
}

func Log(el js.Value) {
	js.Global().Get("window").Get("console").Call("log", el)
}

func StrOrData(args []js.Value, dataAttribute string) string {
	if args[0].Type() == js.TypeString {
		return args[0].String()
	} else {
		return args[0].Call("getAttribute", "data-"+dataAttribute).String()
	}
}

func GetScrollY() int {
	return Win.Get("scrollY").Int()
}

func OverlayScrollTop(div js.Value) {
	div.Get("style").Set("top", strconv.Itoa(GetScrollY())+"px")
}

func OverlayScrollBottom(div js.Value) {
	h := div.Get("offsetHeight").Int()
	posy := GetScrollY() + Win.Get("innerHeight").Int() - h
	go func() {
		time.Sleep(time.Second / 10)
		h2 := div.Get("offsetHeight").Int()
		if h2 != h {
			posy += h
			posy -= h2
			div.Get("style").Set("top", strconv.Itoa(posy)+"px")
		}
	}()
	div.Get("style").Set("top", strconv.Itoa(posy)+"px")
}

func GetValue(inputElement js.Value) string {
	if inputElement.IsNull() || inputElement.IsUndefined() {
		return ""
	}
	val := inputElement.Get("value")
	if val.Type() != js.TypeString {
		return ""
	}
	return val.String()
}

func Confirm(txt string) bool {
	return Win.Call("confirm", txt).Bool()
}

func JsonParse(txt string) js.Value {
	return Win.Get("JSON").Call("parse", txt)
}
