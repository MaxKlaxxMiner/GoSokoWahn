package crc64

type ICrc64 interface {
	UpdateCrc64(currentCrc Value) Value
}
