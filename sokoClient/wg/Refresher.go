package wg

import (
	"syscall/js"
)

type RefreshFunc struct {
	Func             func()
	IdleMilli        int64
	LastRefreshMilli int64
}

// RefreshFuncs Liste mit den aufzurufenen Refresh-Funktionen
var RefreshFuncs []RefreshFunc
var refreshKnown bool

func initRefresher() {
	Wg.Set("refresh", js.FuncOf(func(_ js.Value, args []js.Value) any {
		refreshKnown = false
		performanceTs := int64(args[0].Float())
		for i := range RefreshFuncs {
			if RefreshFuncs[i].LastRefreshMilli+RefreshFuncs[i].IdleMilli > performanceTs {
				continue
			}
			RefreshFuncs[i].Func()
			RefreshFuncs[i].LastRefreshMilli = performanceTs
		}
		return nil
	}))
}

// AddRefreshFunc fügt eine neue Refresh-Funktion hinzu (mit optionalem Idle in Millisekunden)
func AddRefreshFunc(idleMilli int64, f func()) {
	RefreshFuncs = append(RefreshFuncs, RefreshFunc{
		Func:             f,
		IdleMilli:        idleMilli,
		LastRefreshMilli: 0,
	})
}

// NextRefresh gibt an, dass beim nächsten Bildschirm-Update sofort die RefreshFuncs aufgerufen werden soll
func NextRefresh() {
	if refreshKnown {
		return
	}
	Wg.Set("refreshNow", true)
	refreshKnown = true
}
