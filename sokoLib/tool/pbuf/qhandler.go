package pbuf

type Qhandler interface {
	Qappend(buf []byte) []byte
	Qread(buf []byte) int
}
