package pbuf

import (
	"unsafe"
)

func AppendString(buf []byte, val string) []byte {
	buf = AppendVarInt(buf, uint64(len(val)))
	buf = append(buf, val...)
	return buf
}

func WriteString(buf []byte, ofs int, val string) (byteCount int) {
	byteCount = WriteVarInt(buf, ofs, uint64(len(val)))
	ofs += byteCount
	byteCount += len(val)
	copy(buf[ofs:ofs+len(val)], val)
	return
}

func ReadString(buf []byte, ofs int, val *string) (byteCount int) {
	n, byteCount := ConsumeVarInt(buf[ofs:])
	*val = string(buf[ofs+byteCount:][:n])
	byteCount += int(n)
	return byteCount
}

func unsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ReadStringUnsafe(buf []byte, ofs int) (val string, byteCount int) {
	n, byteCount := ConsumeVarInt(buf[ofs:])
	val = unsafeBytesToString(buf[ofs+byteCount:][:n])
	byteCount += int(n)
	return
}

func SizeString(val string) int {
	return SizeVarInt(uint64(len(val))) + len(val)
}
