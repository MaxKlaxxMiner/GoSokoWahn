package pbuf

import (
	"database/sql"
	"math"
	"math/bits"
	"time"
)

func SqlAppendBool(buf []byte, val bool) []byte {
	return AppendBool(buf, val)
}

func SqlAppendNullBool(buf []byte, val sql.NullBool) []byte {
	if val.Valid {
		return append(buf, BoolTo[byte](val.Bool)+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendByte(buf []byte, val byte) []byte {
	return append(buf, val)
}

func SqlAppendNullByte(buf []byte, val sql.NullByte) []byte {
	if val.Valid {
		return AppendVarInt(buf, uint64(val.Byte)+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendInt16(buf []byte, val int16) []byte {
	return AppendVarInt(buf, ZigZag(int64(val)))
}

func SqlAppendUint16(buf []byte, val uint16) []byte {
	return AppendVarInt(buf, uint64(val))
}

func SqlAppendNullInt16(buf []byte, val sql.NullInt16) []byte {
	if val.Valid {
		return AppendVarInt(buf, ZigZag(int64(val.Int16))+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendInt32(buf []byte, val int32) []byte {
	return AppendVarInt(buf, ZigZag(int64(val)))
}

func SqlAppendUint32(buf []byte, val uint32) []byte {
	return AppendVarInt(buf, uint64(val))
}

func SqlAppendNullInt32(buf []byte, val sql.NullInt32) []byte {
	if val.Valid {
		return AppendVarInt(buf, ZigZag(int64(val.Int32))+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendInt64(buf []byte, val int64) []byte {
	return AppendVarInt(buf, ZigZag(val))
}

func SqlAppendUint64(buf []byte, val uint64) []byte {
	return AppendVarInt(buf, val)
}

func SqlAppendInt(buf []byte, val int) []byte {
	return AppendVarInt(buf, ZigZag(int64(val)))
}

func SqlAppendUint(buf []byte, val uint) []byte {
	return AppendVarInt(buf, uint64(val))
}

func SqlAppendNullInt64(buf []byte, val sql.NullInt64) []byte {
	if val.Valid {
		v := ZigZag(val.Int64)
		if v > math.MaxUint64-2 {
			buf = AppendVarInt(buf, math.MaxUint64) // special condition
			v += 2
		}
		return AppendVarInt(buf, v+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendFloat64(buf []byte, val float64) []byte {
	return AppendVarInt(buf, bits.ReverseBytes64(math.Float64bits(val)))
}

func SqlAppendNullFloat64(buf []byte, val sql.NullFloat64) []byte {
	if val.Valid {
		v := bits.ReverseBytes64(math.Float64bits(val.Float64))
		if v > math.MaxUint64-2 {
			buf = AppendVarInt(buf, math.MaxUint64) // special condition
			v += 2
		}
		return AppendVarInt(buf, v+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendString(buf []byte, val string) []byte {
	buf = AppendVarInt(buf, uint64(len(val)))
	buf = append(buf, val...)
	return buf
}

func SqlAppendNullString(buf []byte, val sql.NullString) []byte {
	if val.Valid {
		buf = AppendVarInt(buf, uint64(len(val.String))+1)
		return append(buf, val.String...)
	}
	return append(buf, 0) // Null
}

func SqlAppendTime(buf []byte, val time.Time) []byte {
	return AppendVarInt(buf, uint64(val.UnixMilli()))
}

func SqlAppendNullTime(buf []byte, val sql.NullTime) []byte {
	if val.Valid {
		return AppendVarInt(buf, uint64(val.Time.UnixMilli())+1)
	}
	return append(buf, 0) // Null
}

func SqlAppendBlob(buf []byte, val []byte) []byte {
	return AppendBytes(buf, val)
}
