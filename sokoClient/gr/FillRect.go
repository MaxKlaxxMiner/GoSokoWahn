package gr

func (g *Gr) FillRect(x, y, w, h int) {
	g.Ctx.Call("fillRect", x, y, w, h)
}

func (g *Gr) FillRectS(x, y, w, h int, fillStyle string) {
	g.SetFillStyle(fillStyle)
	g.FillRect(x, y, w, h)
}

func (g *Gr) FillRectF(x, y, w, h float64) {
	g.Ctx.Call("fillRect", x, y, w, h)
}

func (g *Gr) FillRectSF(x, y, w, h float64, fillStyle string) {
	g.SetFillStyle(fillStyle)
	g.FillRectF(x, y, w, h)
}

func (g *Gr) BulkFillRects(rects ...Rect) {
	if len(rects) > 0 {
		win.Call("_bulkFillRects", g.Ctx, sliceToFloat64Array(rects))
	}
}
