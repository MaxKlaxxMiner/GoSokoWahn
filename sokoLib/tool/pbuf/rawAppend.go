package pbuf

import (
	"math"
)

func RawAppendInt8(buf []byte, val int8) []byte {
	buf = append(buf, byte(val))
	return buf
}

func RawAppendInt16(buf []byte, val int16) []byte {
	buf = append(buf, byte(uint16(val)))
	buf = append(buf, byte(uint16(val)>>8))
	return buf
}

func RawAppendInt24(buf []byte, val int32) []byte {
	buf = append(buf, byte(uint32(val)))
	buf = append(buf, byte(uint32(val)>>8))
	buf = append(buf, byte(uint32(val)>>16))
	return buf
}

func RawAppendInt32(buf []byte, val int32) []byte {
	buf = append(buf, byte(uint32(val)))
	buf = append(buf, byte(uint32(val)>>8))
	buf = append(buf, byte(uint32(val)>>16))
	buf = append(buf, byte(uint32(val)>>24))
	return buf
}

func RawAppendInt40(buf []byte, val int64) []byte {
	buf = append(buf, byte(uint64(val)))
	buf = append(buf, byte(uint64(val)>>8))
	buf = append(buf, byte(uint64(val)>>16))
	buf = append(buf, byte(uint64(val)>>24))
	buf = append(buf, byte(uint64(val)>>32))
	return buf
}

func RawAppendInt48(buf []byte, val int64) []byte {
	buf = append(buf, byte(uint64(val)))
	buf = append(buf, byte(uint64(val)>>8))
	buf = append(buf, byte(uint64(val)>>16))
	buf = append(buf, byte(uint64(val)>>24))
	buf = append(buf, byte(uint64(val)>>32))
	buf = append(buf, byte(uint64(val)>>40))
	return buf
}

func RawAppendInt56(buf []byte, val int64) []byte {
	buf = append(buf, byte(uint64(val)))
	buf = append(buf, byte(uint64(val)>>8))
	buf = append(buf, byte(uint64(val)>>16))
	buf = append(buf, byte(uint64(val)>>24))
	buf = append(buf, byte(uint64(val)>>32))
	buf = append(buf, byte(uint64(val)>>40))
	buf = append(buf, byte(uint64(val)>>48))
	return buf
}

func RawAppendInt64(buf []byte, val int64) []byte {
	buf = append(buf, byte(uint64(val)))
	buf = append(buf, byte(uint64(val)>>8))
	buf = append(buf, byte(uint64(val)>>16))
	buf = append(buf, byte(uint64(val)>>24))
	buf = append(buf, byte(uint64(val)>>32))
	buf = append(buf, byte(uint64(val)>>40))
	buf = append(buf, byte(uint64(val)>>48))
	buf = append(buf, byte(uint64(val)>>56))
	return buf
}

func RawAppendUint8(buf []byte, val uint8) []byte {
	buf = append(buf, val)
	return buf
}

func RawAppendUint16(buf []byte, val uint16) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	return buf
}

func RawAppendUint24(buf []byte, val uint32) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	return buf
}

func RawAppendUint32(buf []byte, val uint32) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	buf = append(buf, byte(val>>24))
	return buf
}

func RawAppendUint40(buf []byte, val uint64) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	buf = append(buf, byte(val>>24))
	buf = append(buf, byte(val>>32))
	return buf
}

func RawAppendUint48(buf []byte, val uint64) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	buf = append(buf, byte(val>>24))
	buf = append(buf, byte(val>>32))
	buf = append(buf, byte(val>>40))
	return buf
}

func RawAppendUint56(buf []byte, val uint64) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	buf = append(buf, byte(val>>24))
	buf = append(buf, byte(val>>32))
	buf = append(buf, byte(val>>40))
	buf = append(buf, byte(val>>48))
	return buf
}

func RawAppendUint64(buf []byte, val uint64) []byte {
	buf = append(buf, byte(val))
	buf = append(buf, byte(val>>8))
	buf = append(buf, byte(val>>16))
	buf = append(buf, byte(val>>24))
	buf = append(buf, byte(val>>32))
	buf = append(buf, byte(val>>40))
	buf = append(buf, byte(val>>48))
	buf = append(buf, byte(val>>56))
	return buf
}

func RawAppendInt(buf []byte, val int) []byte {
	return RawAppendInt64(buf, int64(val))
}

func RawAppendUint(buf []byte, val uint) []byte {
	return RawAppendUint64(buf, uint64(val))
}

func RawAppendFloat32(buf []byte, val float32) []byte {
	return RawAppendUint32(buf, math.Float32bits(val))
}

func RawAppendFloat64(buf []byte, val float64) []byte {
	return RawAppendUint64(buf, math.Float64bits(val))
}

func RawAppendFloat64Lite(buf []byte, val float64) []byte {
	return RawAppendFloat32(buf, float32(val))
}

func RawAppendBool(buf []byte, val bool) []byte {
	return AppendBool(buf, val)
}

func RawAppendComplex64(buf []byte, val complex64) []byte {
	buf = RawAppendFloat32(buf, real(val))
	buf = RawAppendFloat32(buf, imag(val))
	return buf
}

func RawAppendComplex128(buf []byte, val complex128) []byte {
	buf = RawAppendFloat64(buf, real(val))
	buf = RawAppendFloat64(buf, imag(val))
	return buf
}
