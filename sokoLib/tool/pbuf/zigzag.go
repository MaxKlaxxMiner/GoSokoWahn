package pbuf

func ZigZag(val int64) uint64 {
	return uint64(val<<1 ^ val>>63)
}

func UnZigZag(val uint64) int64 {
	return int64(val>>1) ^ -int64(val&1)
}
