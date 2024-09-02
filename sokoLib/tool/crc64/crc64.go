package crc64

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/pbuf"
	"math"
	"reflect"
	"time"
	"unsafe"
)

type Value uint64

const (
	Start Value = 0xcbf29ce484222325
	Mul   Value = 0x100000001b3
)

func (crc Value) UpdateZero() Value {
	return crc * Mul
}

func (crc Value) UpdateZero64() Value {
	return crc.UpdateZero().UpdateZero()
}

func (crc Value) UpdateBool(value bool) Value {
	if value {
		return (crc ^ 1) * Mul
	}
	return crc * Mul
}

func (crc Value) UpdateUInt8(value uint8) Value {
	return (crc ^ Value(value)) * Mul
}

func (crc Value) UpdateUInt16(value uint16) Value {
	return (crc ^ Value(value)) * Mul
}

func (crc Value) UpdateUInt32(value uint32) Value {
	return (crc ^ Value(value)) * Mul
}

func (crc Value) UpdateUInt64(value uint64) Value {
	return crc.UpdateUInt32(uint32(value)).UpdateUInt32(uint32(value >> 32))
}

func (crc Value) UpdateUInt(value uint) Value {
	return crc.UpdateUInt64(uint64(value))
}

func (crc Value) UpdateInt8(value int8) Value {
	return crc.UpdateUInt8(uint8(value))
}

func (crc Value) UpdateInt16(value int16) Value {
	return crc.UpdateUInt16(uint16(value))
}

func (crc Value) UpdateInt32(value int32) Value {
	return crc.UpdateUInt32(uint32(value))
}

func (crc Value) UpdateInt64(value int64) Value {
	return crc.UpdateUInt64(uint64(value))
}

func (crc Value) UpdateFloat32(value float32) Value {
	return crc.UpdateUInt32(math.Float32bits(value))
}

func (crc Value) UpdateFloat64(value float64) Value {
	return crc.UpdateUInt64(math.Float64bits(value))
}

func (crc Value) UpdateInt(value int) Value {
	return crc.UpdateUInt(uint(value))
}

func (crc Value) UpdateComplex64(value complex64) Value {
	return crc.UpdateFloat32(real(value)).UpdateFloat32(imag(value))
}

func (crc Value) UpdateComplex128(value complex128) Value {
	return crc.UpdateFloat64(real(value)).UpdateFloat64(imag(value))
}

func (crc Value) UpdateTime(value time.Time) Value {
	return crc.UpdateInt64(value.UnixNano())
}

func updateUnsafeByteSliceAsUInt32Slice(crc Value, data []byte) Value {
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	uintHeader := reflect.SliceHeader{Data: byteHeader.Data, Len: byteHeader.Len >> 2, Cap: byteHeader.Len >> 2}
	uint32Slice := *(*[]uint32)(unsafe.Pointer(&uintHeader))

	result := crc
	for _, val := range uint32Slice {
		result = result.UpdateUInt32ByteOrdered(val)
	}
	return result
}

func (crc Value) UpdateBytes(value []byte) Value {
	result := crc.UpdateUInt32(uint32(len(value)))

	if len(value) >= 4 {
		result = updateUnsafeByteSliceAsUInt32Slice(result, value)
		return result.UpdatePartialBytes(value[uint(len(value))&(math.MaxUint-3):]) // last 0-3 bytes
	}

	return result.UpdatePartialBytes(value)
}

func (crc Value) UpdateString(value string) Value {
	return crc.UpdateBytes(pbuf.UnsafeStringToBytes(value))
}

func (crc Value) UpdatePartialBytes(value []byte) Value {
	result := crc
	for i := 0; i < len(value); i++ {
		result = (result ^ Value(value[i])) * Mul
	}
	return result
}

func (crc Value) UpdatePartialString(value string) Value {
	return crc.UpdatePartialBytes(pbuf.UnsafeStringToBytes(value))
}
