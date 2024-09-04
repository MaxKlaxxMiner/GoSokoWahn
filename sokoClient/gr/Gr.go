package gr

import (
	"sokoClient/wg"
	"strconv"
	"syscall/js"
)

type Gr struct {
	Width        int
	Height       int
	Canvas       js.Value
	Ctx          js.Value
	fillStyle    string
	strokeStyle  string
	lineWidth    float64
	font         string
	textAlign    string
	textBaseline string
}

func (g *Gr) Init(el js.Value) {
	g.Canvas = wg.Doc.Call("createElement", "canvas")
	el.Call("appendChild", g.Canvas)
	g.Ctx = g.Canvas.Call("getContext", "2d")
	g.Resize(0, 0)
}

func (g *Gr) Resize(width, height int) {
	g.Width = width
	g.Height = height
	if g.Width < 1 {
		g.Width = 1
	}
	if g.Height < 1 {
		g.Height = 1
	}
	if width < 1 || height < 1 {
		g.Canvas.Get("style").Set("display", "none")
	} else {
		g.Canvas.Get("style").Set("display", "block")
	}
	g.Canvas.Get("style").Set("width", strconv.Itoa(g.Width)+"px")
	g.Canvas.Get("style").Set("height", strconv.Itoa(g.Height)+"px")
	g.Canvas.Set("width", g.Width)
	g.Canvas.Set("height", g.Height)

	// --- ungültig gewordene Cache-Elemente resetten ---
	g.fillStyle = ""
	g.strokeStyle = ""
	g.lineWidth = 0
	g.font = ""
	g.textAlign = ""
	g.textBaseline = ""
}

func (g *Gr) Clear() {
	g.Ctx.Call("clearRect", 0, 0, g.Width, g.Height)
}
