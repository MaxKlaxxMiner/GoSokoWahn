package gui

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool"
	"sokoClient/gr"
	"sokoClient/wg"
	"syscall/js"
)

type RoomsDebug struct {
	Gui
	main *tdCtx
}

func NewRoomsDebug(container js.Value) *RoomsDebug {
	r := new(RoomsDebug)
	r.InitGui(container)

	r.main = new(tdCtx)

	r.main.frontChanged = func() bool {
		return r.mouseUpdated
	}
	r.main.frontDraw = func(g *gr.Gr) {
		g.SetFont("13px -apple-system, BlinkMacSystemFont, 'Trebuchet MS', Roboto, Ubuntu, sans-serif")
		g.SetTextBaseline("middle")

		g.FillRectS(r.mouseX, r.mouseY, -16, -16, "#0c08")
		g.FillTextS(r.mouseX+8, r.mouseY-8, "active compensation", "#ccc")

		g.FillRectS(r.evIntern.mouseXreal, r.evIntern.mouseYreal-24, -16, -16, "#0c08")
		g.FillTextS(r.evIntern.mouseXreal+8, r.evIntern.mouseYreal-24-8, "not compensated", "#ccc")
	}
	r.AppendRow(r.main)

	wg.AddRefreshFunc(0, func() {
		width := tool.Max(1, container.Get("clientWidth").Int())
		height := tool.Max(1, container.Get("clientHeight").Int())
		if width != r.main.width || height != r.main.height {
			r.main.resize(width, height)
			r.Refresh(true)
		} else {
			r.Refresh(false)
		}
	})

	return r
}
