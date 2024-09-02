package pbuf

var digitmul = []float64{1, 10.00000000001, 100.0000000001, 1000.000000001, 10000.00000001, 100000.0000001, 1000000.000001, 10000000.00001, 100000000.0001, 1000000000.001, 10000000000.01, 100000000000.1}
var digitdiv = []float64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000}

func FloatToInt(val float64, digits int) int {
	return int(val * digitmul[digits])
}

func FloatToIntRounded(val float64, digits int) int {
	if val < 0 {
		return int(val*digitmul[digits] - 0.5)
	} else {
		return int(val*digitmul[digits] + 0.5)
	}
}

func FloatFromInt(val int, digits int) float64 {
	return float64(val) / digitdiv[digits]
}

func AppendFloatInt(buf []byte, val float64, digits int) []byte {
	v := FloatToInt(val, digits)
	return SqlAppendInt(buf, v)
}

func AppendFloatIntRounded(buf []byte, val float64, digits int) []byte {
	v := FloatToIntRounded(val, digits)
	return SqlAppendInt(buf, v)
}

func ReadFloatInt(buf []byte, ofs int, val *float64, digits int) int {
	var v int
	p := SqlReadInt(buf, ofs, &v)
	*val = FloatFromInt(v, digits)
	return p
}

func NewFloatIntDiffer(digits int) *ExtDiffer[float64] {
	return NewExtDiffer(func(last, next float64) (diff float64) {
		return FloatFromInt(FloatToInt(next, digits)-FloatToInt(last, digits), digits)
	}, func(last, diff float64) (next float64) {
		return FloatFromInt(FloatToInt(last, digits)+FloatToInt(diff, digits), digits)
	}, func(buf []byte, val float64) []byte {
		return AppendFloatInt(buf, val, digits)
	}, func(buf []byte, ofs int, val *float64) int {
		return ReadFloatInt(buf, ofs, val, digits)
	}, 0)
}

func NewFloatIntDifferRounded(digits int) *ExtDiffer[float64] {
	return NewExtDiffer(func(last, next float64) (diff float64) {
		return FloatFromInt(FloatToIntRounded(next, digits)-FloatToIntRounded(last, digits), digits)
	}, func(last, diff float64) (next float64) {
		return FloatFromInt(FloatToIntRounded(last, digits)+FloatToIntRounded(diff, digits), digits)
	}, func(buf []byte, val float64) []byte {
		return AppendFloatIntRounded(buf, val, digits)
	}, func(buf []byte, ofs int, val *float64) int {
		return ReadFloatInt(buf, ofs, val, digits)
	}, 0)
}
