package pbuf

type ExtChanger[T comparable] struct {
	qappend     func(buf []byte, val T) []byte
	qread       func(buf []byte, ofs int, val *T) int
	lastChanged bool
	LastValue   T
	startValue  T
}

func NewExtChanger[T comparable](qappend func(buf []byte, val T) []byte, qread func(buf []byte, ofs int, val *T) int, startValue T) *ExtChanger[T] {
	return &ExtChanger[T]{
		qappend:     qappend,
		qread:       qread,
		lastChanged: false,
		LastValue:   startValue,
		startValue:  startValue,
	}
}

func NewExtChangerInt() *ExtChanger[int] {
	return NewExtChanger(SqlAppendInt, SqlReadInt, 0)
}

func NewExtChangerInt16() *ExtChanger[int16] {
	return NewExtChanger(SqlAppendInt16, SqlReadInt16, 0)
}

func NewExtChangerUint() *ExtChanger[uint] {
	return NewExtChanger(SqlAppendUint, SqlReadUint, 0)
}

func NewExtChangerByte() *ExtChanger[byte] {
	return NewExtChanger(SqlAppendByte, SqlReadByte, 0)
}

func NewExtChangerString() *ExtChanger[string] {
	return NewExtChanger(SqlAppendString, SqlReadString, "")
}

func NewExtChangerDictString(dict *ExtDict[string]) *ExtChanger[string] {
	if dict == nil {
		panic("dict == nil")
	}
	return NewExtChanger(func(buf []byte, val string) []byte {
		return ExtAppendDictString(buf, dict, val)
	}, func(buf []byte, ofs int, val *string) int {
		return ExtReadDictString(buf, ofs, dict, val)
	}, "")
}

func (ec *ExtChanger[T]) Clone() *ExtChanger[T] {
	return &ExtChanger[T]{
		qappend:     ec.qappend,
		qread:       ec.qread,
		lastChanged: false,
		LastValue:   ec.startValue,
		startValue:  ec.startValue,
	}
}

func (ec *ExtChanger[T]) Reset() {
	ec.lastChanged = false
	ec.LastValue = ec.startValue
}

func (ec *ExtChanger[T]) ResetValue(val T) {
	ec.lastChanged = false
	ec.LastValue = val
}

func (ec *ExtChanger[T]) PrepareNext(val T) (changed bool) {
	ec.lastChanged = ec.LastValue != val
	ec.LastValue = val
	return ec.lastChanged
}

func (ec *ExtChanger[T]) AppendNext(buf []byte) []byte {
	if !ec.lastChanged {
		return buf
	}
	ec.lastChanged = false
	return ec.qappend(buf, ec.LastValue)
}

func (ec *ExtChanger[T]) AppendNextForce(buf []byte) []byte {
	ec.lastChanged = false
	return ec.qappend(buf, ec.LastValue)
}

func (ec *ExtChanger[T]) ReadNext(buf []byte, ofs int, changed bool, val *T) int {
	ec.lastChanged = changed
	if !changed {
		*val = ec.LastValue
		return 0
	}
	p := ec.qread(buf, ofs, val)
	ec.LastValue = *val
	return p
}
