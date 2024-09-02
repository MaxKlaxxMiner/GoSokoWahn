package deltap

import (
	"errors"
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/pbuf"
)

// Apply returns target by patching origin with delta
func Apply(origin, delta []byte) ([]byte, int, error) {
	return ApplyUseBuf(origin, delta, nil)
}

func ApplyUseBuf(origin, delta, buf []byte) ([]byte, int, error) {
	lenSrc := uint64(len(origin))
	lenDelta := uint64(len(delta))

	limit, bytes := pbuf.ConsumeVarInt(delta)
	if bytes == 0 {
		return nil, 0, errors.New("read error")
	}
	p := uint64(bytes)

	buf = buf[:0]
	if cap(buf) < int(limit) {
		buf = make([]byte, 0, limit)
	}

	for len(buf) < int(limit) {
		var cnt, ofst uint64
		cnt, bytes = pbuf.ConsumeVarInt(delta[p:])
		if bytes == 0 {
			return nil, 0, errors.New("read error")
		}
		p += uint64(bytes)

		if bytes >= len(delta) {
			return nil, 0, errors.New("generated size does not match predicted size")
		}
		if cnt&1 == 0 {
			cnt >>= 1
			if len(buf)+int(cnt) > int(limit) {
				return nil, 0, errors.New("insert command gives an output larger than predicted")
			}
			if cnt > lenDelta {
				return nil, 0, errors.New("insert count exceeds size of delta")
			}
			buf = append(buf, delta[p:p+cnt]...)
			p += cnt
		} else {
			cnt >>= 1
			ofst, bytes = pbuf.ConsumeVarInt(delta[p:])
			if bytes == 0 {
				return nil, 0, errors.New("read error")
			}
			ofst = uint64(int64(len(buf)) - pbuf.UnZigZag(ofst))
			p += uint64(bytes)
			if uint64(len(buf))+cnt > limit {
				return nil, 0, errors.New("copy exceeds output file size")
			}
			if ofst+cnt > lenSrc {
				return nil, 0, errors.New("copy extends past end of input")
			}
			buf = append(buf, origin[ofst:ofst+cnt]...)
		}
	}
	if len(buf) == int(limit) {
		return buf, int(p), nil
	}
	return nil, 0, errors.New("unterminated delta")
}
