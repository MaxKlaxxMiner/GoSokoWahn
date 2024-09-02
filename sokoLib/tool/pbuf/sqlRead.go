package pbuf

import (
	"database/sql"
	"math"
	"math/bits"
	"time"
)

func SqlReadBool(buf []byte, ofs int, val *bool) int {
	return ReadBool(buf, ofs, val)
}

func SqlReadNullBool(buf []byte, ofs int, val *sql.NullBool) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		*val = sql.NullBool{Valid: true, Bool: IntToBool(v - 1)}
	} else {
		*val = sql.NullBool{}
	}
	return b
}

func SqlReadByte(buf []byte, ofs int, val *byte) int {
	*val = buf[ofs]
	return 1
}

func SqlReadNullByte(buf []byte, ofs int, val *sql.NullByte) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		*val = sql.NullByte{Valid: true, Byte: byte(v - 1)}
	} else {
		*val = sql.NullByte{}
	}
	return b
}

func SqlReadInt16(buf []byte, ofs int, val *int16) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = int16(UnZigZag(v))
	return b
}

func SqlReadUint16(buf []byte, ofs int, val *uint16) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = uint16(v)
	return b
}

func SqlReadNullInt16(buf []byte, ofs int, val *sql.NullInt16) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		*val = sql.NullInt16{Valid: true, Int16: int16(UnZigZag(v - 1))}
	} else {
		*val = sql.NullInt16{}
	}
	return b
}

func SqlReadInt32(buf []byte, ofs int, val *int32) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = int32(UnZigZag(v))
	return b
}

func SqlReadUint32(buf []byte, ofs int, val *uint32) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = uint32(v)
	return b
}

func SqlReadNullInt32(buf []byte, ofs int, val *sql.NullInt32) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		*val = sql.NullInt32{Valid: true, Int32: int32(UnZigZag(v - 1))}
	} else {
		*val = sql.NullInt32{}
	}
	return b
}

func SqlReadInt64(buf []byte, ofs int, val *int64) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = UnZigZag(v)
	return b
}

func SqlReadUint64(buf []byte, ofs int, val *uint64) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = v
	return b
}

func SqlReadInt(buf []byte, ofs int, val *int) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = int(int64(UnZigZag(v)))
	return b
}

func SqlReadUint(buf []byte, ofs int, val *uint) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = uint(v)
	return b
}

func SqlReadNullInt64(buf []byte, ofs int, val *sql.NullInt64) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		if v == math.MaxUint64 { // special condition
			b += ReadVarInt(buf, ofs+b, &v)
			v -= 2
		}
		*val = sql.NullInt64{Valid: true, Int64: UnZigZag(v - 1)}
	} else {
		*val = sql.NullInt64{}
	}
	return b
}

func SqlReadFloat64(buf []byte, ofs int, val *float64) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = math.Float64frombits(bits.ReverseBytes64(v))
	return b
}

func SqlReadNullFloat64(buf []byte, ofs int, val *sql.NullFloat64) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		if v == math.MaxUint64 { // special condition
			b += ReadVarInt(buf, ofs+b, &v)
			v -= 2
		}
		*val = sql.NullFloat64{Valid: true, Float64: math.Float64frombits(bits.ReverseBytes64(v - 1))}
	} else {
		*val = sql.NullFloat64{}
	}
	return b
}

func SqlReadString(buf []byte, ofs int, val *string) int {
	return ReadString(buf, ofs, val)
}

func SqlReadNullString(buf []byte, ofs int, val *sql.NullString) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		v--
		*val = sql.NullString{Valid: true, String: string(buf[ofs+b:][:v])}
		b += int(v)
	} else {
		*val = sql.NullString{}
	}
	return b
}

func SqlReadTime(buf []byte, ofs int, val *time.Time) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	*val = time.UnixMilli(int64(v)).UTC()
	return b
}

func SqlReadNullTime(buf []byte, ofs int, val *sql.NullTime) int {
	var v uint64
	b := ReadVarInt(buf, ofs, &v)
	if v != 0 {
		*val = sql.NullTime{Valid: true, Time: time.UnixMilli(int64(v - 1)).UTC()}
	} else {
		*val = sql.NullTime{}
	}
	return b
}

func SqlReadBlob(buf []byte, ofs int, val *[]byte) int {
	bytes, byteCount := ReadBytes(buf, ofs)
	*val = bytes
	return byteCount
}
