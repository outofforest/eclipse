package eclipse

import (
	"math"

	"github.com/pkg/errors"
)

const (
	numOfBits = 7
	maxZeros  = math.MaxUint8 >> 1
)

type encoder struct {
	labels    []byte
	byteIndex uint64
	bitIndex  uint64
}

func newEncoder() *encoder {
	return &encoder{
		labels: []byte{},
	}
}

func (e *encoder) Labels() []byte {
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
			e.labels = append(e.labels, uint8(zeros))
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

func (e *encoder) SetByte(byteIndex uint64, b byte) {
	if numOfBits*byteIndex < e.bitIndex {
		panic(errors.New("index is in the past"))
	}
	if b <= oneByteNumFlag {
		panic(errors.New("invalid byte"))
	}

	if byteIndex > e.byteIndex {
		byteDiff := byteIndex - e.byteIndex
		e.byteIndex += byteDiff
		for byteDiff > 0 {
			zeros := min(byteDiff, maxZeros)
			e.labels = append(e.labels, uint8(zeros))
			byteDiff -= zeros
		}
	}
	e.labels = append(e.labels, b)
	e.byteIndex++
	e.bitIndex = numOfBits * (byteIndex + 1)
}
