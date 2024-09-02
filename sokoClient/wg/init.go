package wg

import (
	"syscall/js"
)

var Wg js.Value

var Win js.Value
var Doc js.Value
var Loc js.Value
var Body js.Value

var UserAgent string
var Storage js.Value
var HashStr string

func Init() {
	g := js.Global()
	Win = g.Get("window")
	Wg = Win.Get("wg")
	Doc = Win.Get("document")
	Loc = Win.Get("location")
	Body = Doc.Get("body")

	UserAgent = Win.Get("navigator").Get("userAgent").String()
	Storage = Win.Get("localStorage")
	HashStr = Loc.Get("hash").String()

	initRefresher()
}
