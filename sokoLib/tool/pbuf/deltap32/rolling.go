package deltap32

type rollingHash struct {
	a, b, i uint16
	z       [nHashSize]byte
}

func newRollingHash() *rollingHash {
	return &rollingHash{}
}

func (rHash *rollingHash) Init(z []byte, pos int) {
	var a, b, x, i uint16

	for i = 0; i < nHashSize; i++ {
		x = uint16(z[pos+int(i)])
		a += x
		b += (nHashSize - i) * x
		rHash.z[i] = byte(x)
	}
	rHash.a = a
	rHash.b = b
	rHash.i = 0
}

func (rHash *rollingHash) Next(c byte) {
	old := uint16(rHash.z[rHash.i])
	rHash.a = rHash.a - old + uint16(c)
	rHash.b = rHash.b - nHashSize*old + rHash.a
	rHash.z[rHash.i] = c
	rHash.i = (rHash.i + 1) & (nHashSize - 1)
}

func (rHash *rollingHash) Value() uint32 {
	return uint32(rHash.a) | (uint32(rHash.b) << 16)
}
