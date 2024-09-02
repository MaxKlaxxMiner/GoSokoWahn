package pbuf

import (
	"time"
)

type KeyTypes interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | time.Duration | float32 | float64 | string
}
