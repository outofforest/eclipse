package eclipse

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"testing"
)

const (
	numOfRecords = 2000000
	density      = 1
)

func BenchmarkReference8(b *testing.B) {
	l1 := [numOfRecords / 8]uint8{}
	l2 := [numOfRecords / 8]uint8{}

	_, _ = rand.Read(l1[:])
	_, _ = rand.Read(l2[:])

	var r uint8
	b.ResetTimer()
	for range b.N {
		for i, l := range l1 {
			r = l & l2[i]
		}
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, r)
}

func BenchmarkReference64(b *testing.B) {
	l1 := [numOfRecords / 64]uint64{}
	l2 := [numOfRecords / 64]uint64{}

	for i := range cap(l1) {
		l1[i] = mathrand.Uint64()
		l2[i] = mathrand.Uint64()
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

func BenchmarkCombined1(b *testing.B) {
	l1 := createLabels()
	l2 := createLabels()
	l3 := createLabels()
	l4 := createLabels()
	l5 := createLabels()
	l6 := createLabels()
	l7 := createLabels()

	var r []uint64

	b.ResetTimer()
	for range b.N {
		r = AND(
			AND(
				OR(l1, l2, l3),
				OR(l4, l5),
			),
			OR(l6, l7),
		)()
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, r)
}

func BenchmarkCombined2(b *testing.B) {
	l1 := createLabels()
	l2 := createLabels()
	l3 := createLabels()
	l4 := createLabels()
	l5 := createLabels()
	l6 := createLabels()
	l7 := createLabels()

	var r []uint64

	b.ResetTimer()
	for range b.N {
		l14 := AND(l1, l4)
		l15 := AND(l1, l5)
		l24 := AND(l2, l4)
		l25 := AND(l2, l5)
		l34 := AND(l3, l4)
		l35 := AND(l3, l5)

		r = OR(
			AND(l14, l6),
			AND(l14, l7),
			AND(l15, l6),
			AND(l15, l7),
			AND(l24, l6),
			AND(l24, l7),
			AND(l25, l6),
			AND(l25, l7),
			AND(l34, l6),
			AND(l34, l7),
			AND(l35, l6),
			AND(l35, l7),
		)()
	}
	b.StopTimer()

	_, _ = fmt.Fprint(io.Discard, r)
}

func BenchmarkLabelEncoding(b *testing.B) {
	indices := make([]bool, numOfRecords)
	for i := range cap(indices) {
		if mathrand.Intn(100) < density {
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

func createLabels() LabelSource {
	e := newEncoder()
	for i := range uint64(numOfRecords) {
		if mathrand.Intn(100) < density {
			e.SetBit(i)
		}
	}
	return Direct(e.Labels())
}
