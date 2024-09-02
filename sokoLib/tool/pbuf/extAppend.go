package pbuf

const debugPanicUnknownDictKey = true

func ExtAppendPrice(buf []byte, val float64) []byte {
	return ExtAppendPriceCents(buf, int(val*100.00000001))
}

func ExtAppendPriceCents(buf []byte, val int) []byte {
	return SqlAppendInt(buf, val)
}

func ExtAppendDualVarInt(buf []byte, high, low uint64, lowBits int) []byte {
	sLowBits := 64 - lowBits
	low = low<<sLowBits>>sLowBits | high<<lowBits
	high >>= sLowBits
	return AppendVarInt128(buf, high, low)
}

const tsOffset int64 = 1704067200000 // 2024-01-01 00:00:00 UTC

func ExtAppendTs(buf []byte, ts int64) []byte { // z.B. 1718718495887
	// marker: 00 = full int64
	// marker: 01 = ofs, millis, fix 5 bytes
	// marker: 10 = ofs, seconds, fix 4 bytes
	// marker: 11 = ofs, minutes, fix 3 bytes

	ts -= tsOffset

	if ts < 0 { // past timestamp
		ts += tsOffset
		return ExtAppendDualVarInt(buf, ZigZag(ts), 0b00, 2)
	}
	if ts%60000 == 0 && ts/60000 < 1<<(24-2) {
		return RawAppendUint24(buf, uint32(ts/60000)<<2|0b11) // minutes, fix 3 bytes
	}
	if ts%1000 == 0 && ts/1000 < 1<<(32-2) {
		return RawAppendUint32(buf, uint32(ts/1000)<<2|0b10) // seconds, fix 4 bytes
	}
	if ts < 1<<(40-2) {
		return RawAppendUint40(buf, uint64(ts)<<2|0b01) // millis, fix 5 bytes
	}

	ts += tsOffset
	return ExtAppendDualVarInt(buf, ZigZag(ts), 0b00, 2)
}

func extFindByte(bytes []byte, searchByte byte) uint64 {
	for i := uint64(0); i < uint64(len(bytes)); i++ {
		if bytes[i] == searchByte {
			return i
		}
	}
	return uint64(len(bytes))
}

func ExtAppendRleString(buf []byte, val string) []byte {
	comp := make([]byte, 0, 64)
	known := make([]byte, 0, 16)

	for i := 0; i < len(val); {
		chr := val[i]
		count := 1
		for i+count < len(val) && val[i+count] == chr {
			count++
		}
		find := extFindByte(known, chr)
		if find == uint64(len(known)) {
			known = append(known, chr)
			find += uint64(len(known)) * uint64(count)
			comp = AppendVarInt(comp, find)
			comp = append(comp, chr)
		} else {
			find += uint64(len(known)+1) * uint64(count)
			comp = AppendVarInt(comp, find)
		}
		if len(comp) >= len(val) {
			break
		}
		i += count
	}

	if len(comp) < len(val) {
		buf = SqlAppendInt64(buf, -int64(len(val)))
		buf = append(buf, comp...)
	} else {
		buf = SqlAppendInt64(buf, int64(len(val)))
		buf = append(buf, val...)
	}
	return buf
}

func ExtAppendBits(buf []byte, bits ...bool) []byte {
	for i, bit := range bits {
		switch i & 7 {
		case 0:
			buf = append(buf, BoolTo[byte](bit))
		case 1:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 1
		case 2:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 2
		case 3:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 3
		case 4:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 4
		case 5:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 5
		case 6:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 6
		case 7:
			buf[len(buf)-1] |= BoolTo[byte](bit) << 7
		}
	}
	return buf
}

func AppendItems[T any](buf []byte, appendItem func([]byte, T) []byte, items ...T) []byte {
	if appendItem == nil {
		return buf
	}
	for i := range items {
		buf = appendItem(buf, items[i])
	}
	return buf
}

func AppendItemsPtr[T any](buf []byte, appendItem func([]byte, *T) []byte, items ...T) []byte {
	if appendItem == nil {
		return buf
	}
	for i := range items {
		buf = appendItem(buf, &items[i])
	}
	return buf
}
