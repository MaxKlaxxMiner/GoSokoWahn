package pbuf

func BoolTo[T uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | uint | int](val bool) T {
	if val {
		return 1
	}
	return 0
}

func BoolToInt(val bool) int {
	return BoolTo[int](val)
}

func IntToBool[T uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | uint | int](val T) bool {
	return val != 0
}

func AppendBool(buf []byte, val bool) []byte {
	return append(buf, BoolTo[byte](val))
}

func ReadBool(buf []byte, ofs int, val *bool) (byteCount int) {
	*val = IntToBool(buf[ofs])
	return 1
}
