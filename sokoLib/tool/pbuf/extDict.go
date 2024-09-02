package pbuf

type ExtDict[T comparable] struct {
	SearchMap map[T]uint64
	UseMap    []T
}

func NewExtDict[T comparable](values ...T) (result *ExtDict[T]) {
	result = &ExtDict[T]{
		SearchMap: make(map[T]uint64, len(values)),
		UseMap:    make([]T, 0, len(values)),
	}
	for _, v := range values {
		result.Add(v)
	}
	return result
}

func (dict *ExtDict[T]) Add(val T) uint64 {
	if r, ok := dict.SearchMap[val]; ok {
		return r
	}
	newId := uint64(len(dict.UseMap))
	dict.SearchMap[val] = newId
	dict.UseMap = append(dict.UseMap, val)
	return newId
}

func ExtAppendDictString(buf []byte, dict *ExtDict[string], val string) []byte {
	if id, ok := dict.SearchMap[val]; ok {
		return AppendVarInt(buf, id)
	} else {
		buf = AppendVarInt(buf, uint64(len(dict.UseMap)+len(val)))
		buf = append(buf, val...)
		if debugPanicUnknownDictKey {
			panic("unknown key: " + val)
		}
		return buf
	}
}

func ExtAppendDictAsRawString(buf []byte, dict *ExtDict[string], val string) []byte {
	buf = AppendVarInt(buf, uint64(len(dict.UseMap)+len(val)))
	buf = append(buf, val...)
	return buf
}

func ExtReadDictString(buf []byte, ofs int, dict *ExtDict[string], val *string) int {
	var v uint64
	p := ReadVarInt(buf, ofs, &v)
	if v < uint64(len(dict.UseMap)) {
		*val = dict.UseMap[v]
	} else {
		v -= uint64(len(dict.UseMap))
		*val = string(buf[ofs+p:][:v])
		p += int(v)
	}
	return p
}

func ExtAppendDynDictString(buf []byte, dict *ExtDict[string], val string) []byte {
	if id, ok := dict.SearchMap[val]; ok {
		return AppendVarInt(buf, id)
	} else {
		buf = AppendVarInt(buf, uint64(len(dict.UseMap)+len(val)))
		buf = append(buf, val...)
		dict.Add(val)
		return buf
	}
}

func ExtReadDynDictString(buf []byte, ofs int, dict *ExtDict[string], val *string) int {
	var v uint64
	p := ReadVarInt(buf, ofs, &v)
	if v < uint64(len(dict.UseMap)) {
		*val = dict.UseMap[v]
	} else {
		v -= uint64(len(dict.UseMap))
		*val = string(buf[ofs+p:][:v])
		p += int(v)
		dict.Add(*val)
	}
	return p
}

func ExtAppendDynDictString1k(buf []byte, dict *ExtDict[string], val string) []byte {
	if id, ok := dict.SearchMap[val]; ok {
		return AppendVarInt1k(buf, id)
	} else {
		buf = AppendVarInt1k(buf, uint64(len(dict.UseMap)+len(val)))
		buf = append(buf, val...)
		dict.Add(val)
		return buf
	}
}

func ExtReadDynDictString1k(buf []byte, ofs int, dict *ExtDict[string], val *string) int {
	var v uint64
	p := ReadVarInt1k(buf, ofs, &v)
	if v < uint64(len(dict.UseMap)) {
		*val = dict.UseMap[v]
	} else {
		v -= uint64(len(dict.UseMap))
		*val = string(buf[ofs+p:][:v])
		p += int(v)
		dict.Add(*val)
	}
	return p
}

func ExtAppendDynDictUint64(buf []byte, dict *ExtDict[uint64], val uint64) []byte {
	if id, ok := dict.SearchMap[val]; ok {
		return AppendVarInt(buf, id)
	} else {
		buf = AppendVarInt(buf, uint64(len(dict.UseMap)))
		buf = RawAppendUint64(buf, val)
		dict.Add(val)
		return buf
	}
}

func ExtReadDynDictUint64(buf []byte, ofs int, dict *ExtDict[uint64], val *uint64) int {
	var v uint64
	p := ReadVarInt(buf, ofs, &v)
	if v < uint64(len(dict.UseMap)) {
		*val = dict.UseMap[v]
	} else {
		p += RawReadUint64(buf, p, val)
		dict.Add(*val)
	}
	return p
}
