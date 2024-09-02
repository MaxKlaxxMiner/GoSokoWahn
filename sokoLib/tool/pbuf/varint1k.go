package pbuf

func AppendVarInt1k(buf []byte, val uint64) []byte {
	if val < 250 { // 0-249
		return append(buf, byte(val))
	}
	if val < 16843770 {
		if val < 506 { // 250-505
			return append(buf, 250, byte(val-250))
		}
		if val < 762 { // 506-761
			return append(buf, 251, byte(val-506))
		}
		if val < 1018 { // 762-1017
			return append(buf, 252, byte(val-762))
		}
		if val < 66554 { // 1018-66553
			val -= 1018
			return append(buf, 253, byte(val), byte(val>>8))
		}
		// 66554-16843769
		val -= 66554
		return append(buf, 254, byte(val), byte(val>>8), byte(val>>16))
	}
	// >= 16843770
	return AppendVarInt(append(buf, 255), val-16843770)
}

func ReadVarInt1k(buf []byte, ofs int, val *uint64) (byteCount int) {
	v := uint64(buf[ofs])
	if v < 250 {
		*val = v
		return 1
	}
	switch v {
	case 250:
		*val = 250 + uint64(buf[ofs+1])
		return 2
	case 251:
		*val = 506 + uint64(buf[ofs+1])
		return 2
	case 252:
		*val = 762 + uint64(buf[ofs+1])
		return 2
	case 253:
		*val = 1018 + uint64(buf[ofs+1]) + (uint64(buf[ofs+2]) << 8)
		return 3
	case 254:
		*val = 66554 + uint64(buf[ofs+1]) + (uint64(buf[ofs+2]) << 8) + (uint64(buf[ofs+3]) << 16)
		return 4
	default:
		v, byteCount = ConsumeVarInt(buf[ofs+1:])
		*val = v + 16843770
		return byteCount + 1
	}
}

func SqlAppendInt1k(buf []byte, val int) []byte {
	return AppendVarInt1k(buf, ZigZag(int64(val)))
}

func SqlAppendUint1k(buf []byte, val uint) []byte {
	return AppendVarInt1k(buf, uint64(val))
}

func SqlReadInt1k(buf []byte, ofs int, val *int) int {
	var v uint64
	b := ReadVarInt1k(buf, ofs, &v)
	*val = int(int64(UnZigZag(v)))
	return b
}

func SqlReadUint1k(buf []byte, ofs int, val *uint) int {
	var v uint64
	b := ReadVarInt1k(buf, ofs, &v)
	*val = uint(v)
	return b
}
