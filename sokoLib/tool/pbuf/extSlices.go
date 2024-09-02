package pbuf

import (
	"time"
)

func ExtAppendSlice[T KeyTypes](buf []byte, val []T) []byte {
	buf = AppendVarInt(buf, uint64(len(val)))
	switch v := any(val).(type) {
	case []int:
		for i := range v {
			buf = SqlAppendInt(buf, v[i])
		}
	case []uint:
		for i := range v {
			buf = SqlAppendUint(buf, v[i])
		}
	case []int8:
		for i := range v {
			buf = append(buf, byte(v[i]))
		}
	case []uint8:
		buf = append(buf, v...)
	case []int16:
		for i := range v {
			buf = SqlAppendInt16(buf, v[i])
		}
	case []uint16:
		for i := range v {
			buf = AppendVarInt(buf, uint64(v[i]))
		}
	case []int32:
		for i := range v {
			buf = SqlAppendInt32(buf, v[i])
		}
	case []uint32:
		for i := range v {
			buf = AppendVarInt(buf, uint64(v[i]))
		}
	case []int64:
		for i := range v {
			buf = SqlAppendInt64(buf, v[i])
		}
	case []uint64:
		for i := range v {
			buf = AppendVarInt(buf, v[i])
		}
	case []time.Duration:
		for i := range v {
			buf = SqlAppendInt64(buf, int64(v[i]))
		}
	case []float32:
		for i := range v {
			buf = RawAppendFloat32(buf, v[i])
		}
	case []float64:
		for i := range v {
			buf = SqlAppendFloat64(buf, v[i])
		}
	case []string:
		for i := range v {
			buf = AppendString(buf, v[i])
		}
	default:
		panic("unsupported type?")
	}
	return buf
}

func ExtReadSlice[T KeyTypes](buf []byte, ofs int, val *[]T) int {
	buf = buf[ofs:]
	var c uint64
	p := ReadVarInt(buf, 0, &c)
	if c == 0 {
		*val = nil
		return p
	}
	r := make([]T, c)
	switch v := any(r).(type) {
	case []int:
		for i := range v {
			p += SqlReadInt(buf, p, &v[i])
		}
	case []uint:
		for i := range v {
			p += SqlReadUint(buf, p, &v[i])
		}
	case []int8:
		for i := range v {
			v[i] = int8(buf[p+i])
		}
		p += len(v)
	case []uint8:
		copy(v, buf[p:])
		p += len(v)
	case []int16:
		for i := range v {
			p += SqlReadInt16(buf, p, &v[i])
		}
	case []uint16:
		var tmp uint64
		for i := range v {
			p += ReadVarInt(buf, p, &tmp)
			v[i] = uint16(tmp)
		}
	case []int32:
		for i := range v {
			p += SqlReadInt32(buf, p, &v[i])
		}
	case []uint32:
		var tmp uint64
		for i := range v {
			p += ReadVarInt(buf, p, &tmp)
			v[i] = uint32(tmp)
		}
	case []int64:
		for i := range v {
			p += SqlReadInt64(buf, p, &v[i])
		}
	case []uint64:
		for i := range v {
			p += ReadVarInt(buf, p, &v[i])
		}
	case []time.Duration:
		var tmp int64
		for i := range v {
			p += SqlReadInt64(buf, p, &tmp)
			v[i] = time.Duration(tmp)
		}
	case []float32:
		for i := range v {
			p += RawReadFloat32(buf, p, &v[i])
		}
	case []float64:
		for i := range v {
			p += SqlReadFloat64(buf, p, &v[i])
		}
	case []string:
		for i := range v {
			p += ReadString(buf, p, &v[i])
		}
	default:
		panic("unsupported type?")
	}
	*val = r
	return p
}

func ExtAppendSliceFunc[T any](buf []byte, val []T, appendFunc func(buf []byte, val *T) []byte) []byte {
	buf = AppendVarInt(buf, uint64(len(val)))
	for i := range val {
		buf = appendFunc(buf, &val[i])
	}
	return buf
}

func ExtReadSliceFunc[T any](buf []byte, ofs int, val *[]T, readFunc func(buf []byte, ofs int, val *T) int) int {
	buf = buf[ofs:]
	var v uint64
	p := ReadVarInt(buf, 0, &v)
	if v == 0 {
		*val = nil
		return p
	}
	r := make([]T, v)
	for i := range r {
		p += readFunc(buf, p, &r[i])
	}
	*val = r
	return p
}
