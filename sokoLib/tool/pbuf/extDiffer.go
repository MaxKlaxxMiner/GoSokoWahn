package pbuf

type ExtDiffer[T any] struct {
	encoder    func(last, next T) (diff T)
	decoder    func(last, diff T) (next T)
	qappend    func(buf []byte, val T) []byte
	qread      func(buf []byte, ofs int, val *T) int
	LastValue  T
	startValue T
}

func NewExtDiffer[T any](encoder func(last, next T) (diff T), decoder func(last, diff T) (next T), qappend func(buf []byte, val T) []byte, qread func(buf []byte, ofs int, val *T) int, startValue T) *ExtDiffer[T] {
	return &ExtDiffer[T]{
		encoder:    encoder,
		decoder:    decoder,
		qappend:    qappend,
		qread:      qread,
		LastValue:  startValue,
		startValue: startValue,
	}
}

func NewExtDifferInt() *ExtDiffer[int] {
	return NewExtDiffer(func(last, next int) (diff int) { return next - last }, func(last, diff int) (next int) { return last + diff }, SqlAppendInt, SqlReadInt, 0)
}

func NewExtDifferInt1k() *ExtDiffer[int] {
	return NewExtDiffer(func(last, next int) (diff int) { return next - last }, func(last, diff int) (next int) { return last + diff }, SqlAppendInt1k, SqlReadInt1k, 0)
}

func NewExtDifferUint() *ExtDiffer[uint] {
	return NewExtDiffer(func(last, next uint) (diff uint) { return next - last }, func(last, diff uint) (next uint) { return last + diff }, SqlAppendUint, SqlReadUint, 0)
}

func NewExtDifferUint1k() *ExtDiffer[uint] {
	return NewExtDiffer(func(last, next uint) (diff uint) { return next - last }, func(last, diff uint) (next uint) { return last + diff }, SqlAppendUint1k, SqlReadUint1k, 0)
}

func NewExtDifferUintDec() *ExtDiffer[uint] {
	return NewExtDiffer(func(last, next uint) (diff uint) { return last - next }, func(last, diff uint) (next uint) { return last - diff }, SqlAppendUint, SqlReadUint, 0)
}

func NewExtDifferUintDec1k() *ExtDiffer[uint] {
	return NewExtDiffer(func(last, next uint) (diff uint) { return last - next }, func(last, diff uint) (next uint) { return last - diff }, SqlAppendUint1k, SqlReadUint1k, 0)
}

func NewExtDifferPrice() *ExtDiffer[float64] {
	return NewExtDiffer(func(last, next float64) (diff float64) {
		return float64(int((next-last)*100.00000001)) / 100.0
	}, func(last, diff float64) (next float64) {
		return float64(int((last+diff)*100.00000001)) / 100.0
	}, ExtAppendPrice, ExtReadPrice, 0)
}

func (ed *ExtDiffer[T]) Clone() *ExtDiffer[T] {
	return &ExtDiffer[T]{
		encoder:    ed.encoder,
		decoder:    ed.decoder,
		qappend:    ed.qappend,
		qread:      ed.qread,
		LastValue:  ed.startValue,
		startValue: ed.startValue,
	}
}

func (ed *ExtDiffer[T]) Reset() {
	ed.LastValue = ed.startValue
}

func (ed *ExtDiffer[T]) ResetValue(val T) {
	ed.LastValue = val
}

func (ed *ExtDiffer[T]) AppendNext(buf []byte, val T) []byte {
	diff := ed.encoder(ed.LastValue, val)
	buf = ed.qappend(buf, diff)
	ed.LastValue = val
	return buf
}

func (ed *ExtDiffer[T]) AppendNextDifErr(buf []byte, val T) ([]byte, T) {
	diff := ed.encoder(ed.LastValue, val)
	buf = ed.qappend(buf, diff)
	ed.LastValue = val
	return buf, diff
}

func (ed *ExtDiffer[T]) ReadNext(buf []byte, ofs int, val *T) int {
	var diff T
	p := ed.qread(buf, ofs, &diff)
	ed.LastValue = ed.decoder(ed.LastValue, diff)
	*val = ed.LastValue
	return p
}

func (ed *ExtDiffer[T]) ReadNextDifErr(buf []byte, ofs int, val *T) (int, T) {
	var diff T
	p := ed.qread(buf, ofs, &diff)
	ed.LastValue = ed.decoder(ed.LastValue, diff)
	*val = ed.LastValue
	return p, diff
}
