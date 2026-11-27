package eclipse

import (
	"math"

	"github.com/pkg/errors"
)

const (
	maxZeros       uint64 = math.MaxUint64 >> 1
	oneByteNumFlag uint64 = maxZeros + 1
	numOfBits             = 63
)

type encoder struct {
	labels    []uint64
	byteIndex uint64
	bitIndex  uint64
}

func newEncoder() *encoder {
	return &encoder{
		labels: []uint64{},
	}
}

func (e *encoder) Labels() []uint64 {
	return e.labels
}

func (e *encoder) SetBit(index uint64) {
	if index < e.bitIndex {
		panic(errors.New("index is in the past"))
	}
	byteIndex := index / numOfBits
	bitIndex := index % numOfBits
	if byteIndex > e.byteIndex {
		byteDiff := byteIndex - e.byteIndex
		e.byteIndex += byteDiff
		for byteDiff > 0 {
			zeros := min(byteDiff, maxZeros)
			e.labels = append(e.labels, zeros)
			byteDiff -= zeros
		}
	}
	if byteIndex == e.byteIndex {
		e.labels = append(e.labels, (1<<bitIndex)|oneByteNumFlag)
		e.byteIndex++
	} else {
		e.labels[len(e.labels)-1] |= 1 << bitIndex
	}
	e.bitIndex = index + 1
}

func (e *encoder) SetChunk(byteIndex uint64, v uint64) {
	if numOfBits*byteIndex < e.bitIndex {
		panic(errors.New("index is in the past"))
	}
	if v <= oneByteNumFlag {
		panic(errors.New("invalid byte"))
	}

	if byteIndex > e.byteIndex {
		byteDiff := byteIndex - e.byteIndex
		e.byteIndex += byteDiff
		for byteDiff > 0 {
			zeros := min(byteDiff, maxZeros)
			e.labels = append(e.labels, zeros)
			byteDiff -= zeros
		}
	}
	e.labels = append(e.labels, v)
	e.byteIndex++
	e.bitIndex = numOfBits * (byteIndex + 1)
}
