package gr

import "math"

func (g *Gr) MoveTo(x, y int) {
	g.Ctx.Call("moveTo", x, y)
}

func (g *Gr) MoveToF(x, y float64) {
	g.Ctx.Call("moveTo", x, y)
}

func (g *Gr) MoveToClean(x, y float64) {
	g.Ctx.Call("moveTo", math.Round(x)+0.5, math.Round(y)+0.5)
}

func (g *Gr) LineTo(x, y int) {
	g.Ctx.Call("lineTo", x, y)
}

func (g *Gr) LineToF(x, y float64) {
	g.Ctx.Call("lineTo", x, y)
}

func (g *Gr) LineToClean(x, y float64) {
	g.Ctx.Call("lineTo", math.Round(x)+0.5, math.Round(y)+0.5)
}

func (g *Gr) FullLineCleanS(x1, y1, x2, y2 float64, strokeStyle string) {
	g.BeginPath()
	g.MoveToClean(x1, y1)
	g.LineToClean(x2, y2)
	g.StrokeS(strokeStyle)
}

func (g *Gr) BulkPolyLine(points ...Pos) {
	if len(points) > 0 {
		win.Call("_bulkPolyLine", g.Ctx, sliceToFloat64Array(points))
	}
}

func (g *Gr) BulkPolyLineI(points ...PosI) {
	if len(points) > 0 {
		win.Call("_bulkPolyLine", g.Ctx, sliceToFloat64Array(points))
	}
}

func (g *Gr) BulkLines(lines ...Line) {
	if len(lines) > 0 {
		win.Call("_bulkLines", g.Ctx, sliceToFloat64Array(lines))
	}
}

func (g *Gr) BulkLinesI(lines ...LineI) {
	if len(lines) > 0 {
		win.Call("_bulkLines", g.Ctx, sliceToFloat64Array(lines))
	}
}

func (g *Gr) Arc(x, y, r, startGrad, endGrad int) {
	g.Ctx.Call("arc", x, y, r, math.Pi/180.0*float64(startGrad), math.Pi/180.0*float64(endGrad))
}

func (g *Gr) ArcF(x, y, r, startRad, endRad float64) {
	g.Ctx.Call("arc", x, y, r, startRad, endRad)
}
