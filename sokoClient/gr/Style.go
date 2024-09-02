package gr

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/crc64"
	"strconv"
	"syscall/js"
)

func (g *Gr) SetFillStyle(style string) {
	if style == g.fillStyle {
		return
	}
	g.Ctx.Set("fillStyle", style)
	g.fillStyle = style
}

func (g *Gr) SetStrokeStyle(style string) {
	if style == g.strokeStyle {
		return
	}
	g.Ctx.Set("strokeStyle", style)
	g.strokeStyle = style
}

func (g *Gr) SetFillStyleGradient(gradient *Gradient) {
	if gradient.crc == g.fillStyle {
		return
	}
	g.Ctx.Set("fillStyle", gradient.el)
	g.fillStyle = gradient.crc
}

func (g *Gr) SetStrokeStyleGradient(gradient *Gradient) {
	if gradient.crc == g.fillStyle {
		return
	}
	g.Ctx.Set("strokeStyle", gradient.el)
	g.fillStyle = gradient.crc
}

func (g *Gr) SetLineDash(steps ...int) {
	anys := make([]any, len(steps))
	for i := range steps {
		anys[i] = steps[i]
	}
	g.Ctx.Call("setLineDash", js.ValueOf(anys))
}

func (g *Gr) SetLineWidth(lineWidth float64) {
	if lineWidth == g.lineWidth {
		return
	}
	g.Ctx.Set("lineWidth", lineWidth)
	g.lineWidth = lineWidth
}

type ColorStop struct {
	Stop  float64
	Color string
}

type Gradient struct {
	el  js.Value
	crc string
}

func (g *Gr) CreateLinearGradient(x1, y1, x2, y2 int, colorStops ...ColorStop) *Gradient {
	crc := crc64.Start.UpdateUInt32(uint32(x1)).UpdateUInt32(uint32(y1)).UpdateUInt32(uint32(x2)).UpdateUInt32(uint32(y2))
	el := g.Ctx.Call("createLinearGradient", x1, y1, x2, y2)
	for _, stop := range colorStops {
		crc = crc.UpdateFloat64(stop.Stop).UpdateString(stop.Color)
		el.Call("addColorStop", stop.Stop, stop.Color)
	}
	return &Gradient{el, strconv.FormatUint(uint64(crc), 16)}
}
