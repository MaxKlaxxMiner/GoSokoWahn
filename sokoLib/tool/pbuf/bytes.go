package pbuf

func AppendBytes(buf []byte, val []byte) []byte {
	buf = AppendVarInt(buf, uint64(len(val)))
	buf = append(buf, val...)
	return buf
}

func WriteBytes(buf []byte, ofs int, val []byte) (byteCount int) {
	byteCount = WriteVarInt(buf, ofs, uint64(len(val)))
	copy(buf[ofs+byteCount:], val)
	return byteCount + len(val)
}

func ReadBytes(buf []byte, ofs int) (val []byte, byteCount int) {
	n, byteCount := ConsumeVarInt(buf[ofs:])
	val = make([]byte, n)
	copy(val, buf[ofs+byteCount:])
	byteCount += int(n)
	return
}

func ReadBytesUnsafe(buf []byte, ofs int) (val []byte, byteCount int) {
	n, byteCount := ConsumeVarInt(buf[ofs:])
	val = buf[ofs+byteCount:][:n]
	byteCount += int(n)
	return
}

func SizeBytes(val []byte) int {
	return SizeVarInt(uint64(len(val))) + len(val)
}
