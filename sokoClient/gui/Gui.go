package gui

import (
	"sokoClient/wg"
	"strconv"
	"syscall/js"
)

type Gui struct {
	Events
	container  js.Value
	mainDiv    js.Value
	tableBody  js.Value
	rows       [][]*tdCtx
	guiUpdated bool
}

func (g *Gui) InitGui(container js.Value) {
	g.container = container
	g.mainDiv = wg.CreateElementFromHTML("<div class='gui' ondragstart='return false' ondrop='return false'></div>")
	g.container.Call("appendChild", g.mainDiv)

	table := wg.Doc.Call("createElement", "table")
	g.tableBody = wg.Doc.Call("createElement", "tbody")
	table.Call("appendChild", g.tableBody)
	g.mainDiv.Call("appendChild", table)

	g.initEvents(container)
}

func (g *Gui) AppendRow(tds ...*tdCtx) {
	if len(tds) == 0 {
		panic("tds is empty")
	}

	tr := wg.Doc.Call("createElement", "tr")
	g.tableBody.Call("appendChild", tr)

	for i := range tds {
		if tds[i] == nil {
			panic("tds[" + strconv.Itoa(i) + "] is nil")
		}
		tds[i].init(tr)
	}
	g.rows = append(g.rows, tds)
}

const mouseCompensationFactor = 4

func (g *Gui) Refresh(forceDraw bool) {
	if g == nil {
		return
	}

	if forceDraw && len(g.rows) > 0 {
		ofsMain := g.container.Get("offsetLeft").Int()
		ofsY := g.container.Get("offsetTop").Int()
		for i := range g.rows {
			tds := g.rows[i]
			ofsX := ofsMain
			for _, td := range tds {
				td.ofsX = ofsX
				td.ofsY = ofsY
				ofsX += td.width
			}
			ofsY += tds[0].height
		}
		g.guiUpdated = true
	}

	g.EventsBeginRefresh()

	// --- Zeichenbereiche neu zeichen (je nach Notwendigkeit) ---
	for i := range g.rows {
		for _, td := range g.rows[i] {
			td.Draw(forceDraw)
		}
	}

	// --- End-Werte verarbeiten ---
	g.EventsEndRefresh()
	g.guiUpdated = false
}
