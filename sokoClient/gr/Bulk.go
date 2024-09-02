package gr

import (
	"fmt"
	"runtime"
	"syscall/js"
	"unsafe"
)

func initMethods() js.Value {
	g := js.Global()
	g.Call("eval", `

	function _bulkPolyLine(ctx, floats) {
		ctx.moveTo(floats[0], floats[1]);
		for (let i = 3; i < floats.length; i += 2) {
			ctx.lineTo(floats[i - 1], floats[i]);
		}
	}

	function _bulkLines(ctx, floats) {
		for (let i = 3; i < floats.length; i += 4) {
			ctx.moveTo(floats[i - 3], floats[i - 2]);
			ctx.lineTo(floats[i - 1], floats[i]);
		}
	}

	function _bulkFillRects(ctx, floats) {
		for (let i = 3; i < floats.length; i += 4) {
			ctx.fillRect(floats[i - 3], floats[i - 2], floats[i - 1], floats[i]);
		}
	}

	`)
	return g.Get("window")
}

var win = initMethods()

type Pos struct {
	X, Y float64
}

type Line struct {
	X1, Y1, X2, Y2 float64
}

type Rect struct {
	X, Y, W, H float64
}

func sliceToFloat64Array(s any) js.Value {
	switch s := s.(type) {
	case []float64:
		a := js.Global().Get("Uint8Array").New(len(s) * 8)
		js.CopyBytesToJS(a, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*8))
		runtime.KeepAlive(s)
		return js.Global().Get("Float64Array").New(a.Get("buffer"), a.Get("byteOffset"), len(s))
	case []Pos:
		a := js.Global().Get("Uint8Array").New(len(s) * 16)
		js.CopyBytesToJS(a, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*16))
		runtime.KeepAlive(s)
		return js.Global().Get("Float64Array").New(a.Get("buffer"), a.Get("byteOffset"), len(s)*2)
	case []Line:
		a := js.Global().Get("Uint8Array").New(len(s) * 32)
		js.CopyBytesToJS(a, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*32))
		runtime.KeepAlive(s)
		return js.Global().Get("Float64Array").New(a.Get("buffer"), a.Get("byteOffset"), len(s)*4)
	case []Rect:
		a := js.Global().Get("Uint8Array").New(len(s) * 32)
		js.CopyBytesToJS(a, unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*32))
		runtime.KeepAlive(s)
		return js.Global().Get("Float64Array").New(a.Get("buffer"), a.Get("byteOffset"), len(s)*4)
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}
}
