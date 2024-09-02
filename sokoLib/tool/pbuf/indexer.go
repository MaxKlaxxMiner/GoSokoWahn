package pbuf

import (
	"sort"
)

type IndexerIndex struct {
	Pos uint
	Len uint
}

func (pos IndexerIndex) CutSlice(slice []byte) []byte {
	return slice[pos.Pos : pos.Pos+pos.Len]
}

type Indexer[T KeyTypes] struct {
	index map[T]IndexerIndex
}

func NewIndexer[T KeyTypes]() *Indexer[T] {
	r := &Indexer[T]{
		index: make(map[T]IndexerIndex),
	}
	return r
}

func (idx *Indexer[T]) Add(key T, pos, length int) {
	if idx.index == nil {
		idx.index = make(map[T]IndexerIndex)
	}
	idx.index[key] = IndexerIndex{uint(pos), uint(length)}
}

func (idx *Indexer[T]) AddStartPos(key T, startPos int) {
	if idx.index == nil {
		idx.index = make(map[T]IndexerIndex)
	}
	idx.index[key] = IndexerIndex{uint(startPos), 0}
}

func (idx *Indexer[T]) AddEndPos(key T, endPos int) {
	t := idx.index[key]
	t.Len = uint(endPos) - t.Pos
	idx.index[key] = t
}

func (idx *Indexer[T]) Get(key T) IndexerIndex {
	return idx.index[key]
}

func (idx *Indexer[T]) GetKeys() []T {
	keys := make([]T, 0, len(idx.index))
	for k := range idx.index {
		keys = append(keys, k)
	}
	return keys
}

func (idx *Indexer[T]) GetKeysSortedByPos() []T {
	keys := idx.GetKeys()
	sort.Slice(keys, func(i, j int) bool {
		return idx.index[keys[i]].Pos < idx.index[keys[j]].Pos
	})
	return keys
}

func (idx *Indexer[T]) Qappend(buf []byte) []byte {
	buf = AppendVarInt(buf, uint64(len(idx.index)))
	keys := idx.GetKeysSortedByPos()
	buf = ExtAppendSlice(buf, keys)
	differ := NewExtDifferUint()
	for i := range keys {
		pos := idx.index[keys[i]]
		buf = differ.AppendNext(buf, pos.Pos)
		buf = SqlAppendUint(buf, pos.Len)
		differ.LastValue += pos.Len
	}
	return buf
}

func (idx *Indexer[T]) Qread(buf []byte) int {
	p := 0
	var c uint64
	p += ReadVarInt(buf, p, &c)
	idx.index = make(map[T]IndexerIndex, c)
	var keys []T
	p += ExtReadSlice(buf, p, &keys)
	differ := NewExtDifferUint()
	for i := range keys {
		var pos IndexerIndex
		p += differ.ReadNext(buf, p, &pos.Pos)
		p += SqlReadUint(buf, p, &pos.Len)
		differ.LastValue += pos.Len
		idx.index[keys[i]] = pos
	}
	return p
}
