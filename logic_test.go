package eclipse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOR(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	requireT.Equal(
		[]byte{0b10101010},
		OR(
			[]byte{},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b11010101},
		OR(
			[]byte{0b11010101},
			[]byte{},
		),
	)

	requireT.Equal(
		[]byte{},
		OR(
			[]byte{},
			[]byte{},
		),
	)

	requireT.Equal(
		[]byte{0b10101010},
		OR(
			nil,
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b11010101},
		OR(
			[]byte{0b11010101},
			nil,
		),
	)

	requireT.Equal(
		[]byte{},
		OR(
			nil,
			nil,
		),
	)

	requireT.Equal(
		[]byte{0b11110111},
		OR(
			[]byte{0b11010101},
			[]byte{0b10100111},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101},
		OR(
			[]byte{0b00000010, 0b11010101},
			[]byte{0b10100111},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101},
		OR(
			[]byte{0b10100111},
			[]byte{0b00000010, 0b11010101},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101, 0b10001000},
		OR(
			[]byte{0b00000010, 0b11010101},
			[]byte{0b10100111, 0b00000010, 0b10001000},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101, 0b10001000},
		OR(
			[]byte{0b10100111, 0b00000010, 0b10001000},
			[]byte{0b00000010, 0b11010101},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101, 0b10011100},
		OR(
			[]byte{0b10100111, 0b00000010, 0b10001000},
			[]byte{0b00000010, 0b11010101, 0b10010100},
		),
	)

	requireT.Equal(
		[]byte{0b10100111, 0b00000001, 0b11010101, 0b10011100},
		OR(
			[]byte{0b00000010, 0b11010101, 0b10010100},
			[]byte{0b10100111, 0b00000010, 0b10001000},
		),
	)
}

func TestAND(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{0b11010101},
			[]byte{},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{},
			[]byte{},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			nil,
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{0b11010101},
			nil,
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			nil,
			nil,
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{0b11010101},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{0b00000001, 0b10101010},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{},
		AND(
			[]byte{0b10101010},
			[]byte{0b00000001, 0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b10101010},
		AND(
			[]byte{0b10101010},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b10101010},
		AND(
			[]byte{0b10101010, 0b10000001},
			[]byte{0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b10101010},
		AND(
			[]byte{0b10101010},
			[]byte{0b10101010, 0b10000001},
		),
	)

	requireT.Equal(
		[]byte{0b00000001, 0b10001000},
		AND(
			[]byte{0b00000001, 0b10101010},
			[]byte{0b10101010, 0b10001101},
		),
	)

	requireT.Equal(
		[]byte{0b00000001, 0b10001000},
		AND(
			[]byte{0b10101010, 0b10001101},
			[]byte{0b00000001, 0b10101010},
		),
	)

	requireT.Equal(
		[]byte{0b00000010, 0b10101010, 0b10001101},
		AND(
			[]byte{0b00000010, 0b10101010, 0b11111111},
			[]byte{0b10101010, 0b10000000, 0b11111111, 0b10001101},
		),
	)

	requireT.Equal(
		[]byte{0b00000010, 0b10101010, 0b10001101},
		AND(
			[]byte{0b10101010, 0b10000000, 0b11111111, 0b10001101},
			[]byte{0b00000010, 0b10101010, 0b11111111},
		),
	)

	requireT.Equal(
		[]byte{0b00000010, 0b10101010, 0b10001101, 0b00000010, 0b10001000},
		AND(
			[]byte{0b00000010, 0b10101010, 0b11111111, 0b00000001, 0b10000010, 0b10111000},
			[]byte{0b10101010, 0b10000001, 0b11111111, 0b10001101, 0b11111111, 0b00000001, 0b10001111},
		),
	)

	requireT.Equal(
		[]byte{0b00000010, 0b10101010, 0b10001101, 0b00000010, 0b10001000},
		AND(
			[]byte{0b10101010, 0b10000001, 0b11111111, 0b10001101, 0b11111111, 0b00000001, 0b10001111},
			[]byte{0b00000010, 0b10101010, 0b11111111, 0b00000001, 0b10000010, 0b10111000},
		),
	)

	requireT.Equal(
		[]byte{0b00000011, 0b10001101, 0b00000010, 0b10001000},
		AND(
			[]byte{0b00000010, 0b11110000, 0b11111111, 0b00000001, 0b10000010, 0b10111000},
			[]byte{0b10101010, 0b10000001, 0b10001111, 0b10001101, 0b11111111, 0b00000001, 0b10001111},
		),
	)

	requireT.Equal(
		[]byte{0b00000011, 0b10001101, 0b00000010, 0b10001000},
		AND(
			[]byte{0b10101010, 0b10000001, 0b10001111, 0b10001101, 0b11111111, 0b00000001, 0b10001111},
			[]byte{0b00000010, 0b11110000, 0b11111111, 0b00000001, 0b10000010, 0b10111000},
		),
	)

	requireT.Equal(
		[]byte{0b00000011, 0b10001101},
		AND(
			[]byte{0b00000010, 0b11110000, 0b11111111, 0b00000001, 0b10000010, 0b11111000},
			[]byte{0b10101010, 0b10000001, 0b10001111, 0b10001101, 0b11111111, 0b00000001, 0b10000111},
		),
	)

	requireT.Equal(
		[]byte{0b00000011, 0b10001101},
		AND(
			[]byte{0b10101010, 0b10000001, 0b10001111, 0b10001101, 0b11111111, 0b00000001, 0b10000111},
			[]byte{0b00000010, 0b11110000, 0b11111111, 0b00000001, 0b10000010, 0b11111000},
		),
	)
}
