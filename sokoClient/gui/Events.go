package gui

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool"
	"sokoClient/wg"
	"syscall/js"
)

type Events struct {
	evIntern struct {
		mouseEvent js.Func
		mouseXreal int
		mouseXlast int
		mouseYreal int
		mouseYlast int
		mouseBlast int
	}
	mouseX        int  // aktuelle globale Maus-X Position
	mouseY        int  // aktuelle globale Maus-Y Position
	mouseB        int  // aktuell gedrückte Maustasten (1 = linke, 2 = rechte, 3 = beide)
	mouseS        int  // aktuelle Scroll-Position (~ 100 pro Scroll-Schritt)
	mouseLastX    int  // die letzte Maus-X Position (um eine Änderung beim Refresh zu erkennen)
	mouseLastY    int  // die letzte Maus-Y Position (um eine Änderung beim Refresh zu erkennen)
	mouseLastB    int  // die zuletzt gedrückten Maustasten (um eine Änderung beim Refresh zu erkennen)
	mouseLastS    int  // die letzte Scroll-Position
	mouseClickX   int  // Maus-X Position, wo die Maustaste gedrückt wurde (und ggf. gehalten wird) -> nur gültig, wenn mouseB > 0
	mouseClickY   int  // Maus-Y Position, wo die Maustaste gedrückt wurde (und ggf. gehalten wird) -> nur gültig, wenn mouseB > 0
	mouseClickDbl bool // gibt an, ob gerade ein Doppelklick gemacht wurde
	mouseUpdated  bool // wenn die Maus bewegt/geklickt wurde also Änderungen erkannt wurde (zählt z.B. nicht, wenn eine Maustaste gedrückt gehalten wird, nur beim Klick selbst oder beim weiteren Bewegen)
}

func (e *Events) initEvents(container js.Value) {
	// --- Maus-Events behandeln ---
	e.evIntern.mouseXreal = -1
	e.evIntern.mouseYreal = -1
	e.evIntern.mouseEvent = js.FuncOf(func(_ js.Value, args []js.Value) any {
		e.evIntern.mouseXreal = args[0].Get("x").Int()
		e.evIntern.mouseYreal = args[0].Get("y").Int()
		e.mouseB = args[0].Get("buttons").Int()
		if e.mouseB != e.evIntern.mouseBlast { // direct sync bei Maustasten-Änderung
			if e.evIntern.mouseBlast == 0 {
				e.mouseClickX = e.evIntern.mouseXreal
				e.mouseClickY = e.evIntern.mouseYreal
			}
			e.evIntern.mouseBlast = e.mouseB
			e.mouseX = e.evIntern.mouseXreal
			e.mouseY = e.evIntern.mouseYreal
			e.evIntern.mouseXlast = e.mouseX
			e.evIntern.mouseYlast = e.mouseY
		}
		e.mouseUpdated = true
		wg.NextRefresh()
		return nil
	})
	container.Set("onmousedown", e.evIntern.mouseEvent)
	container.Set("onmousemove", e.evIntern.mouseEvent)
	container.Set("onmouseup", e.evIntern.mouseEvent)
	container.Set("onmouseleave", e.evIntern.mouseEvent)
	container.Set("ondblclick", js.FuncOf(func(_ js.Value, args []js.Value) any {
		e.mouseClickDbl = true
		return e.evIntern.mouseEvent.Invoke(args[0])
	}))
	container.Set("onwheel", js.FuncOf(func(_ js.Value, args []js.Value) any {
		e.mouseS += args[0].Get("deltaY").Int()
		e.mouseUpdated = true
		wg.NextRefresh()
		return nil
	}))
}

func (e *Events) EventsBeginRefresh() {
	// --- Maus-Berechnung und -Kompensation durchführen ---
	if e.mouseB != 0 { // Kompensation bei gedrückter Maustaste ignorieren
		e.mouseUpdated = tool.VarUpdate(&e.mouseX, e.evIntern.mouseXreal) || e.mouseUpdated
		e.mouseUpdated = tool.VarUpdate(&e.mouseY, e.evIntern.mouseYreal) || e.mouseUpdated
		e.evIntern.mouseXlast = e.evIntern.mouseXreal
		e.evIntern.mouseYlast = e.evIntern.mouseYreal
	} else {
		mx := e.evIntern.mouseXreal - e.evIntern.mouseXlast
		my := e.evIntern.mouseYreal - e.evIntern.mouseYlast
		if mx < -mouseCompensationFactor/2 || mx > mouseCompensationFactor/2 {
			mx = mx * mouseCompensationFactor / 2
		}
		if my < -mouseCompensationFactor/2 || my > mouseCompensationFactor/2 {
			my = my * mouseCompensationFactor / 2
		}
		e.mouseUpdated = tool.VarUpdate(&e.mouseX, e.evIntern.mouseXreal+mx) || e.mouseUpdated
		e.mouseUpdated = tool.VarUpdate(&e.mouseY, e.evIntern.mouseYreal+my) || e.mouseUpdated
		e.mouseUpdated = tool.VarUpdate(&e.evIntern.mouseXlast, e.evIntern.mouseXreal) || e.mouseUpdated
		e.mouseUpdated = tool.VarUpdate(&e.evIntern.mouseYlast, e.evIntern.mouseYreal) || e.mouseUpdated
	}
}

func (e *Events) EventsEndRefresh() {
	if e.mouseUpdated {
		wg.NextRefresh() // einen weiteren Refresh unmittelbar danach ausführen
	}
	e.mouseUpdated = false
	e.mouseLastX = e.mouseX
	e.mouseLastY = e.mouseY
	e.mouseLastB = e.mouseB
	e.mouseLastS = e.mouseS
	e.mouseClickDbl = false
}
