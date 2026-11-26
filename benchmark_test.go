package eclipse

import (
	"math/rand"
	"testing"
)

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
	indices := make([]bool, 2000000)
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
	for i := range uint64(2000000) {
		if rand.Intn(100) < 20 {
			e.SetBit(i)
		}
	}
	return e.Labels()
}
