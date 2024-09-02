package pbuf

func ExtReadPrice(buf []byte, ofs int, val *float64) int {
	var v int
	p := ExtReadPriceCents(buf, ofs, &v)
	*val = float64(v) / 100.0
	return p
}

func ExtReadPriceCents(buf []byte, ofs int, val *int) int {
	return SqlReadInt(buf, ofs, val)
}

func ExtReadDualVarInt(buf []byte, ofs int, high, low *uint64, lowBits int) int {
	h, l, byteCount := ConsumeVarInt128(buf[ofs:])
	sLowBits := 64 - lowBits
	*high = h<<sLowBits | l>>lowBits
	*low = l << sLowBits >> sLowBits
	return byteCount
}

func ExtReadTs(buf []byte, ofs int, val *int64) int {
	// marker: 00 = full int64
	// marker: 01 = ofs, millis, fix 5 bytes
	// marker: 10 = ofs, seconds, fix 4 bytes
	// marker: 11 = ofs, minutes, fix 3 bytes

	switch buf[ofs] & 0x3 {
	case 0b01: // ofs, millis, fix 5 bytes
		var v uint64
		p := RawReadUint40(buf, ofs, &v)
		*val = int64(v>>2) + tsOffset
		return p
	case 0b10: // ofs, seconds, fix 4 bytes
		var v uint32
		p := RawReadUint32(buf, ofs, &v)
		*val = int64(v>>2)*1000 + tsOffset
		return p
	case 0b11: // ofs, minutes, fix 3 bytes
		var v uint32
		p := RawReadUint24(buf, ofs, &v)
		*val = int64(v>>2)*60000 + tsOffset
		return p
	default: // 0b00 - full int64
		var h, l uint64
		p := ExtReadDualVarInt(buf, ofs, &h, &l, 2)
		*val = UnZigZag(h)
		return p
	}
}

func ExtReadRleString(buf []byte, ofs int, val *string) int {
	buf = buf[ofs:]
	var v int64
	p := SqlReadInt64(buf, 0, &v)
	if v >= 0 {
		*val = string(buf[p:][:v])
		p += int(v)
	} else {
		known := make([]byte, 0, 16)
		v = -v
		tmp := make([]byte, 0, v)

		for len(tmp) < int(v) {
			var find uint64
			var chr byte
			p += ReadVarInt(buf, p, &find)
			count := int(int32(find) / int32(len(known)+1))
			f := uint(find) - uint(count)*uint(len(known)+1)
			if f < uint(len(known)) {
				chr = known[f]
			} else {
				chr = buf[p]
				p++
				known = append(known, chr)
			}
			for ; count > 0; count-- {
				tmp = append(tmp, chr)
			}
		}

		*val = unsafeBytesToString(tmp)
	}
	return p
}

func ExtReadBits(buf []byte, ofs int, bits ...*bool) int {
	p := 0
	var b byte
	for i, bit := range bits {
		switch i & 7 {
		case 0:
			b = buf[ofs+p]
			p++
			*bit = IntToBool(b & 1)
		case 1:
			*bit = IntToBool(b >> 1 & 1)
		case 2:
			*bit = IntToBool(b >> 2 & 1)
		case 3:
			*bit = IntToBool(b >> 3 & 1)
		case 4:
			*bit = IntToBool(b >> 4 & 1)
		case 5:
			*bit = IntToBool(b >> 5 & 1)
		case 6:
			*bit = IntToBool(b >> 6 & 1)
		case 7:
			*bit = IntToBool(b >> 7 & 1)
		}
	}
	return p
}

func ExtReadBitsSlice(buf []byte, ofs int, bits []bool) int {
	p := 0
	var b byte
	for i := range bits {
		switch i & 7 {
		case 0:
			b = buf[ofs+p]
			p++
			bits[i] = IntToBool(b & 1)
		case 1:
			bits[i] = IntToBool(b >> 1 & 1)
		case 2:
			bits[i] = IntToBool(b >> 2 & 1)
		case 3:
			bits[i] = IntToBool(b >> 3 & 1)
		case 4:
			bits[i] = IntToBool(b >> 4 & 1)
		case 5:
			bits[i] = IntToBool(b >> 5 & 1)
		case 6:
			bits[i] = IntToBool(b >> 6 & 1)
		case 7:
			bits[i] = IntToBool(b >> 7 & 1)
		}
	}
	return p
}
