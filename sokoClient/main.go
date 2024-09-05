package main

import (
	"sokoClient/gui"
	"sokoClient/wg"
)

func main() {
	wg.Init()
	//wg.Win.Set("oncontextmenu", js.FuncOf(func(_ js.Value, args []js.Value) any { return false })) // rechtsklick deaktivieren

	rooms := gui.NewRoomsDebug(wg.Body)
	_ = rooms

	select {} // hold
}
