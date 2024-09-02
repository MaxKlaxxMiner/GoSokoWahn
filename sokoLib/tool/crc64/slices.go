package crc64

func UpdateSliceFunc[T any](crc Value, slice []T, crcFunc func(crc Value, val T) Value) Value {
	crc = crc.UpdateUInt32(uint32(len(slice)))
	for i := range slice {
		crc = crcFunc(crc, slice[i])
	}
	return crc
}

func UpdateSliceFuncPtr[T any](crc Value, slice []T, crcFunc func(crc Value, val *T) Value) Value {
	crc = crc.UpdateUInt32(uint32(len(slice)))
	for i := range slice {
		crc = crcFunc(crc, &slice[i])
	}
	return crc
}
