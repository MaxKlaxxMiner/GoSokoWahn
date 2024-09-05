package gui

import (
	"fmt"
	"sokoClient/gr"
	"sokoClient/wg"
	"strconv"
	"syscall/js"
	"time"
)

const drawDebug = false

type tdCtx struct {
	width        int
	height       int
	ofsX         int
	ofsY         int
	td           js.Value
	back         gr.Gr
	backDraw     func(g *gr.Gr)
	backChanged  func() bool
	front        gr.Gr
	frontDraw    func(g *gr.Gr)
	frontChanged func() bool
	onResize     func()       // optional: wird aufgerufen, wenn die Größe des Fenstern geändert wurde (mindestens jedoch einmal am Start, vor den ersten Zeichnen-Aufgaben)
	bgGradient   *gr.Gradient // optional: Hintergrund-Gradient (wenn gesetzt, wird dieser statt den leeren Hintergrund gezeichnet)
}

func (t *tdCtx) init(tr js.Value) {
	t.td = wg.Doc.Call("createElement", "td")
	t.td.Get("style").Set("position", "relative")
	tr.Call("appendChild", t.td)

	t.back.Init(t.td)
	t.back.Canvas.Get("style").Set("position", "absolute")
	t.back.Canvas.Get("style").Set("top", "0")
	t.back.Canvas.Get("style").Set("left", "0")
	t.back.Canvas.Get("style").Set("zIndex", "1")

	t.front.Init(t.td)
	t.front.Canvas.Get("style").Set("position", "absolute")
	t.front.Canvas.Get("style").Set("top", "0")
	t.front.Canvas.Get("style").Set("left", "0")
	t.front.Canvas.Get("style").Set("zIndex", "2")

	t.resize(0, 0)
}

func (t *tdCtx) resize(width, height int) bool {
	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}
	if t.width == width && t.height == height {
		return false
	}
	t.width = width
	t.height = height

	t.td.Get("style").Set("width", strconv.Itoa(width)+"px")
	t.td.Get("style").Set("height", strconv.Itoa(height)+"px")

	t.front.Resize(width, height)
	t.back.Resize(width, height)

	if t.onResize != nil {
		t.onResize()
	}

	return true
}

func (t *tdCtx) Draw(force bool) bool {
	if t.width*t.height == 0 {
		return false
	}
	if t.backChanged == nil {
		t.backChanged = func() bool { return false }
	}
	if t.backDraw == nil {
		t.backDraw = func(g *gr.Gr) {}
	}
	if t.frontChanged == nil {
		t.frontChanged = func() bool { return false }
	}
	if t.frontDraw == nil {
		t.frontDraw = func(g *gr.Gr) {}
	}

	updateBack := t.backChanged()
	updateFront := t.frontChanged()

	if force {
		updateBack = true
		updateFront = true
	}
	if updateBack {
		if t.bgGradient != nil {
			t.back.SetFillStyleGradient(t.bgGradient)
			t.back.FillRect(0, 0, t.back.Width, t.back.Height)
		} else {
			t.back.SetFillStyle("#0000")
			t.back.Clear()
		}
		t.backDraw(&t.back)
		if drawDebug {
			fmt.Println("drawdbg-back", t.ofsX, t.ofsY, time.Now().UnixMilli())
		}
	}
	if updateFront {
		t.front.Clear()
		t.frontDraw(&t.front)
		if drawDebug {
			fmt.Println("drawdbg-front", t.ofsX, t.ofsY, time.Now().UnixMilli())
		}
	}
	return updateBack || updateFront
}

func (t *tdCtx) CheckXY(x, y int) bool {
	return x >= t.ofsX && x < t.ofsX+t.width && y >= t.ofsY && y < t.ofsY+t.height
}

func (t *tdCtx) CheckXs(xs ...int) bool {
	for _, x := range xs {
		if x >= t.ofsX && x < t.ofsX+t.width {
			return true
		}
	}
	return false
}

func (t *tdCtx) CheckYs(ys ...int) bool {
	for _, y := range ys {
		if y >= t.ofsY && y < t.ofsY+t.height {
			return true
		}
	}
	return false
}
