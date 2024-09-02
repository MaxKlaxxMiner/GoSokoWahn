package deltap8

import (
	"github.com/MaxKlaxxMiner/GoSokoWahn/sokoLib/tool/pbuf"
)

const nHashSize = 8

// Create returns the difference between origin and target
func Create(origin, target []byte, searchLimit int) []byte {
	if searchLimit <= 0 {
		searchLimit = 1024
	}
	zDelta := make([]byte, 0, nHashSize)

	zDelta = pbuf.AppendVarInt(zDelta, uint64(len(target)))

	if len(origin) <= nHashSize {
		zDelta = pbuf.AppendVarInt(zDelta, uint64(len(target))<<1|0)
		zDelta = append(zDelta, target...)
		return zDelta
	}

	nHash := len(origin) / nHashSize
	collide := make([]int32, nHash)
	landmark := make([]int32, nHash)

	for i := 0; i < nHash; i++ {
		collide[i] = -1
		landmark[i] = -1
	}

	h := newRollingHash()
	for i := 0; i < len(origin)-nHashSize; i += nHashSize {
		h.Init(origin, i)
		hv := int(h.Value() % uint32(nHash))
		collide[i/nHashSize] = landmark[hv]
		landmark[hv] = int32(i / nHashSize)
	}

	var _base, iSrc, iBlock, bestCnt, bestOfst, bestLitsz int

	for _base+nHashSize < len(target) {
		bestOfst = 0
		bestLitsz = 0
		h.Init(target, _base)
		bestCnt = 0
		for i := 0; ; {
			limit := searchLimit
			hv := int(h.Value() % uint32(nHash))
			iBlock = int(landmark[hv])
			for iBlock >= 0 {
				var cnt, ofst, litsz int
				var j, k, x, y int

				iSrc = iBlock * nHashSize

				x = iSrc
				y = _base + i
				j = 0
				for x < len(origin) && y < len(target) {
					if origin[x] != target[y] {
						break
					}
					x++
					y++
					j++
				}

				for k = 1; k < iSrc && k <= i; k++ {
					if origin[iSrc-k] != target[_base+i-k] {
						break
					}
				}
				k--

				ofst = iSrc - k
				cnt = j + k
				litsz = i - k

				if cnt > bestCnt && cnt > pbuf.SizeVarInt(uint64(i-k))+pbuf.SizeVarInt(uint64(cnt))+pbuf.SizeVarInt(pbuf.ZigZag(int64(_base-ofst))) {
					bestCnt = cnt
					bestOfst = iSrc - k
					bestLitsz = litsz
				}

				iBlock = int(collide[iBlock])

				limit--
				if limit <= 0 {
					break
				}
			}

			if bestCnt > 0 {
				if bestLitsz > 0 {
					zDelta = pbuf.AppendVarInt(zDelta, uint64(bestLitsz)<<1|0)
					zDelta = append(zDelta, target[_base:_base+bestLitsz]...)
					_base += bestLitsz
				}

				zDelta = pbuf.AppendVarInt(zDelta, uint64(bestCnt)<<1|1)
				zDelta = pbuf.AppendVarInt(zDelta, pbuf.ZigZag(int64(_base)-int64(bestOfst)))
				_base += bestCnt
				bestCnt = 0
				break
			}

			if _base+i+nHashSize >= len(target) {
				zDelta = pbuf.AppendVarInt(zDelta, uint64(len(target)-_base)<<1|0)
				zDelta = append(zDelta, target[_base:]...)
				_base = len(target)
				break
			}

			h.Next(target[_base+i+nHashSize])
			i++
		}
	}

	if _base < len(target) {
		zDelta = pbuf.AppendVarInt(zDelta, uint64(len(target)-_base)<<1|0)
		zDelta = append(zDelta, target[_base:]...)
	}

	return zDelta
}
