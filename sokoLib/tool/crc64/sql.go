package crc64

import (
	"database/sql"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/pbuf"
)

func (crc Value) UpdateNullBool(value sql.NullBool) Value {
	if value.Valid {
		return crc.UpdateUInt8(pbuf.BoolTo[byte](value.Bool) + 1)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullByte(value sql.NullByte) Value {
	if value.Valid {
		return crc.UpdateUInt16(uint16(value.Byte) + 1)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullInt16(value sql.NullInt16) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateInt16(value.Int16)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullInt32(value sql.NullInt32) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateInt32(value.Int32)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullInt64(value sql.NullInt64) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateInt64(value.Int64)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullFloat64(value sql.NullFloat64) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateFloat64(value.Float64)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullString(value sql.NullString) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateString(value.String)
	}
	return crc.UpdateZero() // Null
}

func (crc Value) UpdateNullTime(value sql.NullTime) Value {
	if value.Valid {
		return crc.UpdateUInt8(1).UpdateTime(value.Time)
	}
	return crc.UpdateZero() // Null
}
