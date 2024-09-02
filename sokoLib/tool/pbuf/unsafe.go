package pbuf

import (
	"unsafe"
)

func UnsafeBytesToString(val []byte) string {
	//return *(*string)(unsafe.Pointer(&val))
	if len(val) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(val), len(val))
}

func UnsafeStringToBytes(val string) []byte {
	//strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
	//valHeader := reflect.SliceHeader{Data: strHeader.Data, Len: strHeader.Len, Cap: strHeader.Len}
	//return *(*[]byte)(unsafe.Pointer(&valHeader))
	if val == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(val), len(val))
}
