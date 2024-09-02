package gr

func (g *Gr) BeginPath() {
	g.Ctx.Call("beginPath")
}

func (g *Gr) Fill() {
	g.Ctx.Call("fill")
}

func (g *Gr) FillS(style string) {
	g.SetFillStyle(style)
	g.Fill()
}

func (g *Gr) Stroke() {
	g.Ctx.Call("stroke")
}

func (g *Gr) StrokeS(style string) {
	g.SetStrokeStyle(style)
	g.Stroke()
}
