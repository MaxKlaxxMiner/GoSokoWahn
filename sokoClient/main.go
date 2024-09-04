package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/pbuf"
	"sokoClient/gr"
	"sokoClient/wg"
	"syscall/js"
	"time"
)

func main() {
	wg.Init()

	n := "Hello Client!"

	fmt.Println(n, pbuf.AppendString(nil, n))

	//wg.Win.Set("oncontextmenu", js.FuncOf(func(_ js.Value, args []js.Value) any { return false })) // rechtsklick deaktivieren

	var mouseX, mouseY, mouseB, mouseS int
	var mouseUpdated, mouseClickDbl bool

	mouseEvent := js.FuncOf(func(_ js.Value, args []js.Value) any {
		mouseX = args[0].Get("x").Int()
		mouseY = args[0].Get("y").Int()
		mouseB = args[0].Get("buttons").Int()
		mouseUpdated = true
		wg.NextRefresh()
		return nil
	})
	container := wg.Body

	container.Set("onmousedown", mouseEvent)
	container.Set("onmousemove", mouseEvent)
	container.Set("onmouseup", mouseEvent)
	container.Set("onmouseleave", mouseEvent)
	container.Set("ondblclick", js.FuncOf(func(_ js.Value, args []js.Value) any {
		mouseClickDbl = true
		return mouseEvent.Invoke(args[0])
	}))
	container.Set("onwheel", js.FuncOf(func(_ js.Value, args []js.Value) any {
		mouseS += args[0].Get("deltaY").Int()
		mouseUpdated = true
		wg.NextRefresh()
		return nil
	}))

	g := new(gr.Gr)
	g.Init(container)

	_, _ = mouseUpdated, mouseClickDbl
	_, _, _ = mouseX, mouseY, mouseB

	wg.AddRefreshFunc(0, func() {
		width := tool.Max(1, container.Get("clientWidth").Int())
		height := tool.Max(1, container.Get("clientHeight").Int())
		if width != g.Width || height != g.Height {
			g.Resize(width, height)
		}

		g.Clear()
		g.FillRectS(mouseX, mouseY, -16, -16, "#0c08")

		wg.GetElementById("h2title").Set("innerHTML", fmt.Sprintf("%s - %d, %d - %d, %d", time.Now().Format("2006-01-02 15:04:05.000"), mouseX, mouseY, mouseB, mouseS))
	})

	select {} // hold
}
