package pbuf

import (
	"io"
	"math/bits"
)

func AppendVarInt(buf []byte, val uint64) []byte {
	switch {
	case val < 1<<7:
		buf = append(buf, byte(val))
	case val < 1<<14:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte(val>>7))
	case val < 1<<21:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte(val>>14))
	case val < 1<<28:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte(val>>21))
	case val < 1<<35:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte(val>>28))
	case val < 1<<42:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte((val>>28)&0x7f|0x80),
			byte(val>>35))
	case val < 1<<49:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte((val>>28)&0x7f|0x80),
			byte((val>>35)&0x7f|0x80),
			byte(val>>42))
	case val < 1<<56:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte((val>>28)&0x7f|0x80),
			byte((val>>35)&0x7f|0x80),
			byte((val>>42)&0x7f|0x80),
			byte(val>>49))
	case val < 1<<63:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte((val>>28)&0x7f|0x80),
			byte((val>>35)&0x7f|0x80),
			byte((val>>42)&0x7f|0x80),
			byte((val>>49)&0x7f|0x80),
			byte(val>>56))
	default:
		buf = append(buf,
			byte((val>>0)&0x7f|0x80),
			byte((val>>7)&0x7f|0x80),
			byte((val>>14)&0x7f|0x80),
			byte((val>>21)&0x7f|0x80),
			byte((val>>28)&0x7f|0x80),
			byte((val>>35)&0x7f|0x80),
			byte((val>>42)&0x7f|0x80),
			byte((val>>49)&0x7f|0x80),
			byte((val>>56)&0x7f|0x80),
			1)
	}
	return buf
}

func AppendVarInt128(buf []byte, high, low uint64) []byte {
	if high != 0 {
		for i := 0; i < 9; i++ { // write 63 low-bits
			buf = append(buf, byte(low|0x80))
			low >>= 7
		}
		low |= high << 1
		if high >= 1<<63 {
			buf = append(buf, byte(low|0x80))
			low = low>>7 | 1<<(64-7)
		}
	}
	return AppendVarInt(buf, low)
}

func WriteVarInt(buf []byte, ofs int, val uint64) (byteCount int) {
	return SetVarInt(buf[ofs:], val)
}

func SetVarInt(buf []byte, val uint64) (byteCount int) {
	switch {
	case val < 1<<7:
		if len(buf) < 1 {
			return 0
		}
		buf[0] = byte(val)
		return 1
	case val < 1<<14:
		if len(buf) < 2 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val >> 7)
		return 2
	case val < 1<<21:
		if len(buf) < 3 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val >> 14)
		return 3
	case val < 1<<28:
		if len(buf) < 4 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val >> 21)
		return 4
	case val < 1<<35:
		if len(buf) < 5 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val >> 28)
		return 5
	case val < 1<<42:
		if len(buf) < 6 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val>>28 | 0x80)
		buf[5] = byte(val >> 35)
		return 6
	case val < 1<<49:
		if len(buf) < 7 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val>>28 | 0x80)
		buf[5] = byte(val>>35 | 0x80)
		buf[6] = byte(val >> 42)
		return 7
	case val < 1<<56:
		if len(buf) < 8 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val>>28 | 0x80)
		buf[5] = byte(val>>35 | 0x80)
		buf[6] = byte(val>>42 | 0x80)
		buf[7] = byte(val >> 49)
		return 8
	case val < 1<<63:
		if len(buf) < 9 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val>>28 | 0x80)
		buf[5] = byte(val>>35 | 0x80)
		buf[6] = byte(val>>42 | 0x80)
		buf[7] = byte(val>>49 | 0x80)
		buf[8] = byte(val >> 56)
		return 9
	default:
		if len(buf) < 10 {
			return 0
		}
		buf[0] = byte(val | 0x80)
		buf[1] = byte(val>>7 | 0x80)
		buf[2] = byte(val>>14 | 0x80)
		buf[3] = byte(val>>21 | 0x80)
		buf[4] = byte(val>>28 | 0x80)
		buf[5] = byte(val>>35 | 0x80)
		buf[6] = byte(val>>42 | 0x80)
		buf[7] = byte(val>>49 | 0x80)
		buf[8] = byte(val>>56 | 0x80)
		buf[9] = 1
		return 10
	}
}

func SetVarIntFixSize(buf []byte, val uint64, fixByteCount int) (byteCount int) {
	byteCount = SizeVarInt(val)
	if fixByteCount < byteCount || fixByteCount > 10 {
		return 0
	}
	SetVarInt(buf, val)
	for byteCount < fixByteCount {
		buf[byteCount-1] |= 0x80
		buf[byteCount] = 0
		byteCount++
	}
	return byteCount
}

func ReadVarInt(buf []byte, ofs int, val *uint64) (byteCount int) {
	v, byteCount := ConsumeVarInt(buf[ofs:])
	*val = v
	return byteCount
}

func ReadVarInt128(buf []byte, ofs int, high, low *uint64) (byteCount int) {
	h, l, byteCount := ConsumeVarInt128(buf[ofs:])
	*high = h
	*low = l
	return byteCount
}

func ConsumeVarInt(buf []byte) (val uint64, byteCount int) {
	var y uint64
	if len(buf) <= 0 {
		return 0, 0
	}
	val = uint64(buf[0])
	if val < 0x80 {
		return val, 1
	}
	val -= 0x80

	if len(buf) <= 1 {
		return 0, 0
	}
	y = uint64(buf[1])
	val += y << 7
	if y < 0x80 {
		return val, 2
	}
	val -= 0x80 << 7

	if len(buf) <= 2 {
		return 0, 0
	}
	y = uint64(buf[2])
	val += y << 14
	if y < 0x80 {
		return val, 3
	}
	val -= 0x80 << 14

	if len(buf) <= 3 {
		return 0, 0
	}
	y = uint64(buf[3])
	val += y << 21
	if y < 0x80 {
		return val, 4
	}
	val -= 0x80 << 21

	if len(buf) <= 4 {
		return 0, 0
	}
	y = uint64(buf[4])
	val += y << 28
	if y < 0x80 {
		return val, 5
	}
	val -= 0x80 << 28

	if len(buf) <= 5 {
		return 0, 0
	}
	y = uint64(buf[5])
	val += y << 35
	if y < 0x80 {
		return val, 6
	}
	val -= 0x80 << 35

	if len(buf) <= 6 {
		return 0, 0
	}
	y = uint64(buf[6])
	val += y << 42
	if y < 0x80 {
		return val, 7
	}
	val -= 0x80 << 42

	if len(buf) <= 7 {
		return 0, 0
	}
	y = uint64(buf[7])
	val += y << 49
	if y < 0x80 {
		return val, 8
	}
	val -= 0x80 << 49

	if len(buf) <= 8 {
		return 0, 0
	}
	y = uint64(buf[8])
	val += y << 56
	if y < 0x80 {
		return val, 9
	}
	val -= 0x80 << 56

	if len(buf) <= 9 {
		return 0, 0
	}
	y = uint64(buf[9])
	val += y << 63
	if y < 2 {
		return val, 10
	}
	return 0, 0
}

func ConsumeVarInt128(buf []byte) (high, low uint64, byteCount int) {
	var y uint64
	bit := 0
	for b := 0; b <= 8; b++ {
		if len(buf) <= b {
			return 0, 0, 0
		}
		y = uint64(buf[b])
		low += y << bit
		if y < 0x80 {
			return high, low, b + 1
		}
		low -= 0x80 << bit
		bit += 7
	}

	if len(buf) <= 9 {
		return 0, 0, 0
	}
	y = uint64(buf[9])
	low += y << 63
	high = y >> 1
	if y < 0x80 {
		return high, low, 10
	}
	high -= 0x80 >> 1

	bit = 6
	for b := 10; b <= 17; b++ {
		if len(buf) <= b {
			return 0, 0, 0
		}
		y = uint64(buf[b])
		high += y << bit
		if y < 0x80 {
			return high, low, b + 1
		}
		high -= 0x80 << bit
		bit += 7
	}

	if len(buf) <= 18 {
		return 0, 0, 0
	}
	y = uint64(buf[18])
	high += y << 62
	if y < 4 {
		return high, low, 19
	}
	return 0, 0, 0
}

func SizeVarInt(v uint64) int {
	return int(9*uint32(bits.Len64(v))+64) / 64
}

func InsertVarInt(buf []byte, pos int, val uint64) []byte {
	intLen := SizeVarInt(val)
	for i := 0; i < intLen; i++ {
		buf = append(buf, 0)
	}
	copy(buf[pos+intLen:], buf[pos:])
	SetVarInt(buf[pos:], val)
	return buf
}

func InsertVarIntUsePlaceholder(buf []byte, pos, placeholderBytes int, val uint64) []byte {
	intLen := SizeVarInt(val)
	if intLen != placeholderBytes {
		if intLen > placeholderBytes {
			needBytes := intLen - placeholderBytes
			for ; needBytes != 0; needBytes-- {
				buf = append(buf, 0)
			}
			copy(buf[pos+intLen:], buf[pos+placeholderBytes:])
		} else if intLen < placeholderBytes {
			removeBytes := placeholderBytes - intLen
			copy(buf[pos+intLen:], buf[pos+placeholderBytes:])
			buf = buf[:len(buf)-removeBytes]
		}
	}
	SetVarInt(buf[pos:], val)
	return buf
}

func ReadVarIntFromReader(r io.Reader) (val uint64, byteCount int, err error) {
	byteBuf := []byte{0}

	if _, err = io.ReadFull(r, byteBuf); err != nil {
		return
	}
	byteCount++
	val = uint64(byteBuf[0])
	if val > 0x7f {
		val &= 0x7f
		for bit := 7; ; bit += 7 {
			if _, err = io.ReadFull(r, byteBuf); err != nil {
				break
			}
			byteCount++
			val |= uint64(byteBuf[0]&0x7f) << bit
			if byteBuf[0] < 0x7f {
				break
			}
		}
	}
	return
}
