package pbuf

import "math"

//func ExtReadPrice(buf []byte, ofs int, val *float64) int {

func RawReadInt8(buf []byte, ofs int, val *int8) int {
	*val = int8(buf[ofs])
	return 1
}

func RawReadInt16(buf []byte, ofs int, val *int16) int {
	_ = buf[ofs+1]
	t := uint16(buf[ofs])
	t |= uint16(buf[ofs+1]) << 8
	*val = int16(t)
	return 2
}

func RawReadInt24(buf []byte, ofs int, val *int32) int {
	_ = buf[ofs+2]
	t := uint32(buf[ofs])
	t |= uint32(buf[ofs+1]) << 8
	t |= uint32(buf[ofs+2]) << 16
	*val = int32(t)
	return 3
}

func RawReadInt32(buf []byte, ofs int, val *int32) int {
	_ = buf[ofs+3]
	t := uint32(buf[ofs])
	t |= uint32(buf[ofs+1]) << 8
	t |= uint32(buf[ofs+2]) << 16
	t |= uint32(buf[ofs+3]) << 24
	*val = int32(t)
	return 4
}

func RawReadInt40(buf []byte, ofs int, val *int64) int {
	_ = buf[ofs+4]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	*val = int64(t)
	return 5
}

func RawReadInt48(buf []byte, ofs int, val *int64) int {
	_ = buf[ofs+5]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	*val = int64(t)
	return 6
}

func RawReadInt56(buf []byte, ofs int, val *int64) int {
	_ = buf[ofs+6]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	t |= uint64(buf[ofs+6]) << 48
	*val = int64(t)
	return 7
}

func RawReadInt64(buf []byte, ofs int, val *int64) int {
	_ = buf[ofs+7]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	t |= uint64(buf[ofs+6]) << 48
	t |= uint64(buf[ofs+7]) << 56
	*val = int64(t)
	return 8
}

func RawReadUint8(buf []byte, ofs int, val *uint8) int {
	*val = buf[ofs]
	return 1
}

func RawReadUint16(buf []byte, ofs int, val *uint16) int {
	_ = buf[ofs+1]
	t := uint16(buf[ofs])
	t |= uint16(buf[ofs+1]) << 8
	*val = t
	return 2
}

func RawReadUint24(buf []byte, ofs int, val *uint32) int {
	_ = buf[ofs+2]
	t := uint32(buf[ofs])
	t |= uint32(buf[ofs+1]) << 8
	t |= uint32(buf[ofs+2]) << 16
	*val = t
	return 3
}

func RawReadUint32(buf []byte, ofs int, val *uint32) int {
	_ = buf[ofs+3]
	t := uint32(buf[ofs])
	t |= uint32(buf[ofs+1]) << 8
	t |= uint32(buf[ofs+2]) << 16
	t |= uint32(buf[ofs+3]) << 24
	*val = t
	return 4
}

func RawReadUint40(buf []byte, ofs int, val *uint64) int {
	_ = buf[ofs+4]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	*val = t
	return 5
}

func RawReadUint48(buf []byte, ofs int, val *uint64) int {
	_ = buf[ofs+5]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	*val = t
	return 6
}

func RawReadUint56(buf []byte, ofs int, val *uint64) int {
	_ = buf[ofs+6]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	t |= uint64(buf[ofs+6]) << 48
	*val = t
	return 7
}

func RawReadUint64(buf []byte, ofs int, val *uint64) int {
	_ = buf[ofs+7]
	t := uint64(buf[ofs])
	t |= uint64(buf[ofs+1]) << 8
	t |= uint64(buf[ofs+2]) << 16
	t |= uint64(buf[ofs+3]) << 24
	t |= uint64(buf[ofs+4]) << 32
	t |= uint64(buf[ofs+5]) << 40
	t |= uint64(buf[ofs+6]) << 48
	t |= uint64(buf[ofs+7]) << 56
	*val = t
	return 8
}

func RawReadInt(buf []byte, ofs int, val *int) int {
	var t int64
	p := RawReadInt64(buf, ofs, &t)
	*val = int(t)
	return p
}

func RawReadUint(buf []byte, ofs int, val *uint) int {
	var t uint64
	p := RawReadUint64(buf, ofs, &t)
	*val = uint(t)
	return p
}

func RawReadFloat32(buf []byte, ofs int, val *float32) int {
	var t uint32
	p := RawReadUint32(buf, ofs, &t)
	*val = math.Float32frombits(t)
	return p
}

func RawReadFloat64(buf []byte, ofs int, val *float64) int {
	var t uint64
	p := RawReadUint64(buf, ofs, &t)
	*val = math.Float64frombits(t)
	return p
}

func RawReadFloat64Lite(buf []byte, ofs int, val *float64) int {
	var t uint32
	p := RawReadUint32(buf, ofs, &t)
	*val = float64(math.Float32frombits(t))
	return p
}

func RawReadBool(buf []byte, ofs int, val *bool) int {
	return ReadBool(buf, ofs, val)
}

func RawReadComplex64(buf []byte, ofs int, val *complex64) int {
	var r, i float32
	p := RawReadFloat32(buf, ofs, &r)
	p += RawReadFloat32(buf, ofs+p, &i)
	*val = complex(r, i)
	return p
}

func RawReadComplex128(buf []byte, ofs int, val *complex128) int {
	var r, i float64
	p := RawReadFloat64(buf, ofs, &r)
	p += RawReadFloat64(buf, ofs+p, &i)
	*val = complex(r, i)
	return p
}
