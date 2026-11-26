package eclipse

import (
	"fmt"
	"io"
	"math/rand"
	"testing"
)

const (
	numOfRecords = 2000000
	density      = 20
)

func BenchmarkReference(b *testing.B) {
	l1 := [numOfRecords / 64]uint64{}
	l2 := [numOfRecords / 64]uint64{}

	for i := range cap(l1) {
		l1[i] = rand.Uint64()
		l2[i] = rand.Uint64()
	}

	var r uint64
	b.ResetTimer()
	for range b.N {
		for i, l := range l1 {
			r = l & l2[i]
		}
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, r)
}

func BenchmarkAND(b *testing.B) {
	l1 := createLabels()
	l2 := createLabels()

	b.ResetTimer()
	for range b.N {
		AND(l1, l2)
	}
	b.StopTimer()
}

func BenchmarkOR(b *testing.B) {
	l1 := createLabels()
	l2 := createLabels()

	b.ResetTimer()
	for range b.N {
		AND(l1, l2)
	}
	b.StopTimer()
}

func BenchmarkLabelEncoding(b *testing.B) {
	indices := make([]bool, numOfRecords)
	for i := range cap(indices) {
		if rand.Intn(100) < 20 {
			indices[i] = true
		}
	}

	b.ResetTimer()
	for range b.N {
		e := newEncoder()
		for i, v := range indices {
			if v {
				e.SetBit(uint64(i))
			}
		}
	}
	b.StopTimer()
}

func createLabels() []byte {
	e := newEncoder()
	for i := range uint64(numOfRecords) {
		if rand.Intn(100) < density {
			e.SetBit(i)
		}
	}
	return e.Labels()
}
