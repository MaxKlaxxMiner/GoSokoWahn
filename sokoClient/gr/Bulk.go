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

type PosI struct {
	X, Y int
}

type LineI struct {
	X1, Y1, X2, Y2 int
}

type RectI struct {
	X, Y, W, H int
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
	case []PosI:
		tmp := make([]Pos, len(s))
		for i := 0; i < len(tmp); i++ {
			tmp[i].X = float64(s[i].X)
			tmp[i].Y = float64(s[i].Y)
		}
		return sliceToFloat64Array(tmp)
	case []LineI:
		tmp := make([]Line, len(s))
		for i := 0; i < len(tmp); i++ {
			tmp[i].X1 = float64(s[i].X1)
			tmp[i].Y1 = float64(s[i].Y1)
			tmp[i].X2 = float64(s[i].X2)
			tmp[i].Y2 = float64(s[i].Y2)
		}
		return sliceToFloat64Array(tmp)
	case []RectI:
		tmp := make([]Rect, len(s))
		for i := 0; i < len(tmp); i++ {
			tmp[i].X = float64(s[i].X)
			tmp[i].Y = float64(s[i].Y)
			tmp[i].W = float64(s[i].W)
			tmp[i].H = float64(s[i].H)
		}
		return sliceToFloat64Array(tmp)
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}
}
