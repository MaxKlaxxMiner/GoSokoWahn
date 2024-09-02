package gr

func (g *Gr) FillText(x, y int, txt string) {
	g.Ctx.Call("fillText", txt, x, y)
}

func (g *Gr) FillTextF(x, y float64, txt string) {
	g.Ctx.Call("fillText", txt, x, y)
}

func (g *Gr) FillTextS(x, y int, txt, fillStyle string) {
	g.SetFillStyle(fillStyle)
	g.FillText(x, y, txt)
}

func (g *Gr) FillTextFS(x, y float64, txt, fillStyle string) {
	g.SetFillStyle(fillStyle)
	g.FillTextF(x, y, txt)
}

func (g *Gr) MeasureText(txt string) (width float64) {
	r := g.Ctx.Call("measureText", txt)
	return r.Get("width").Float()
}

func (g *Gr) MeasureTextFull(txt string) (actualBoundingBoxAscent, actualBoundingBoxDescent, actualBoundingBoxLeft, actualBoundingBoxRight, alphabeticBaseline, fontBoundingBoxAscent, fontBoundingBoxDescent, hangingBaseline, ideographicBaseline, width float64) {
	r := g.Ctx.Call("measureText", txt)
	return r.Get("actualBoundingBoxAscent").Float(),
		r.Get("actualBoundingBoxDescent").Float(),
		r.Get("actualBoundingBoxLeft").Float(),
		r.Get("actualBoundingBoxRight").Float(),
		r.Get("alphabeticBaseline").Float(),
		r.Get("fontBoundingBoxAscent").Float(),
		r.Get("fontBoundingBoxDescent").Float(),
		r.Get("hangingBaseline").Float(),
		r.Get("ideographicBaseline").Float(),
		r.Get("width").Float()
}

func (g *Gr) SetFont(font string) {
	if g.font == font {
		return
	}
	g.Ctx.Set("font", font)
	g.font = font
}

// SetTextAlign start, end, left, right, center
func (g *Gr) SetTextAlign(textAlign string) {
	if g.textAlign == textAlign {
		return
	}
	g.Ctx.Set("textAlign", textAlign)
	g.textAlign = textAlign
}

// SetTextBaseline top, hanging, middle, alphabetic, ideographic, bottom
func (g *Gr) SetTextBaseline(textBaseline string) {
	if g.textBaseline == textBaseline {
		return
	}
	g.Ctx.Set("textBaseline", textBaseline)
	g.textBaseline = textBaseline
}
