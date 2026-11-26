package eclipse

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoderEmpty(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()
	requireT.Equal([]byte{}, e.Labels())
}

func TestEncoderSetBitPanics(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(0)
	requireT.Panics(func() {
		e.SetBit(0)
	})
}

func TestEncoderSetBitBeginning(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(0)
	requireT.Equal([]byte{0b10000001}, e.Labels())

	e.SetBit(1)
	requireT.Equal([]byte{0b10000011}, e.Labels())

	e.SetBit(2)
	requireT.Equal([]byte{0b10000111}, e.Labels())

	e.SetBit(3)
	requireT.Equal([]byte{0b10001111}, e.Labels())

	e.SetBit(4)
	requireT.Equal([]byte{0b10011111}, e.Labels())

	e.SetBit(5)
	requireT.Equal([]byte{0b10111111}, e.Labels())

	e.SetBit(6)
	requireT.Equal([]byte{0b11111111}, e.Labels())

	e.SetBit(7)
	requireT.Equal([]byte{0b11111111, 0b10000001}, e.Labels())

	e.SetBit(8)
	requireT.Equal([]byte{0b11111111, 0b10000011}, e.Labels())

	e.SetBit(9)
	requireT.Equal([]byte{0b11111111, 0b10000111}, e.Labels())
}

func TestEncoderSetBitBeginningWithHoles(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(0)
	requireT.Equal([]byte{0b10000001}, e.Labels())

	e.SetBit(2)
	requireT.Equal([]byte{0b10000101}, e.Labels())

	e.SetBit(4)
	requireT.Equal([]byte{0b10010101}, e.Labels())

	e.SetBit(6)
	requireT.Equal([]byte{0b11010101}, e.Labels())

	e.SetBit(8)
	requireT.Equal([]byte{0b11010101, 0b10000010}, e.Labels())
}

func TestEncoderSetBitSecondByte(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(7)
	requireT.Equal([]byte{0b00000001, 0b10000001}, e.Labels())

	e.SetBit(9)
	requireT.Equal([]byte{0b00000001, 0b10000101}, e.Labels())

	e.SetBit(11)
	requireT.Equal([]byte{0b00000001, 0b10010101}, e.Labels())

	e.SetBit(13)
	requireT.Equal([]byte{0b00000001, 0b11010101}, e.Labels())

	e.SetBit(15)
	requireT.Equal([]byte{0b00000001, 0b11010101, 0b10000010}, e.Labels())

	e.SetBit(17)
	requireT.Equal([]byte{0b00000001, 0b11010101, 0b10001010}, e.Labels())
}

func TestEncoderSetBitThirdByte(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(14)
	requireT.Equal([]byte{0b00000010, 0b10000001}, e.Labels())

	e.SetBit(16)
	requireT.Equal([]byte{0b00000010, 0b10000101}, e.Labels())

	e.SetBit(18)
	requireT.Equal([]byte{0b00000010, 0b10010101}, e.Labels())

	e.SetBit(20)
	requireT.Equal([]byte{0b00000010, 0b11010101}, e.Labels())

	e.SetBit(22)
	requireT.Equal([]byte{0b00000010, 0b11010101, 0b10000010}, e.Labels())

	e.SetBit(24)
	requireT.Equal([]byte{0b00000010, 0b11010101, 0b10001010}, e.Labels())
}

func TestEncoderSetBitHoles(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(21)
	requireT.Equal([]byte{0b00000011, 0b10000001}, e.Labels())

	e.SetBit(23)
	requireT.Equal([]byte{0b00000011, 0b10000101}, e.Labels())

	e.SetBit(25)
	requireT.Equal([]byte{0b00000011, 0b10010101}, e.Labels())

	e.SetBit(27)
	requireT.Equal([]byte{0b00000011, 0b11010101}, e.Labels())

	e.SetBit(29)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10000010}, e.Labels())

	e.SetBit(31)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010}, e.Labels())

	e.SetBit(49)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010, 0b00000010, 0b10000001}, e.Labels())

	e.SetBit(50)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010, 0b00000010, 0b10000011}, e.Labels())

	e.SetBit(51)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010, 0b00000010, 0b10000111}, e.Labels())

	e.SetBit(53)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010, 0b00000010, 0b10010111}, e.Labels())

	e.SetBit(57)
	requireT.Equal([]byte{0b00000011, 0b11010101, 0b10001010, 0b00000010, 0b10010111, 0b10000010}, e.Labels())
}

func TestEncoderSetBigHoles(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetBit(889)
	requireT.Equal([]byte{0b01111111, 0b10000001}, e.Labels())

	e.SetBit(1786)
	requireT.Equal([]byte{0b01111111, 0b10000001, 0b01111111, 0b10000010}, e.Labels())

	e.SetBit(3586)
	requireT.Equal(
		[]byte{0b01111111, 0b10000001, 0b01111111, 0b10000010, 0b01111111, 0b01111111, 0b00000010, 0b10000100},
		e.Labels(),
	)
}

func TestEncoderSetBytePanics(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetByte(0, 0b10000001)
	requireT.Panics(func() {
		e.SetByte(0, 0b10000001)
	})
	requireT.Panics(func() {
		e.SetByte(2, 0b10000000)
	})
	requireT.Panics(func() {
		e.SetByte(2, 0b00000001)
	})
	e.SetByte(2, 0b10000001)
}

func TestEncoderSetByteBeginning(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetByte(0, 0b11010101)
	requireT.Equal([]byte{0b11010101}, e.Labels())

	e.SetByte(1, 0b10101010)
	requireT.Equal([]byte{0b11010101, 0b10101010}, e.Labels())

	e.SetByte(2, 0b11111111)
	requireT.Equal([]byte{0b11010101, 0b10101010, 0b11111111}, e.Labels())
}

func TestEncoderSetByteBeginningWithHole(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetByte(1, 0b10101010)
	requireT.Equal([]byte{0b00000001, 0b10101010}, e.Labels())

	e.SetByte(2, 0b11111111)
	requireT.Equal([]byte{0b00000001, 0b10101010, 0b11111111}, e.Labels())
}

func TestEncoderSetByteBeginningWithHoles(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)

	e := newEncoder()

	e.SetByte(254, 0b10101010)
	requireT.Equal([]byte{0b01111111, 0b01111111, 0b10101010}, e.Labels())

	e.SetByte(385, 0b11111111)
	requireT.Equal([]byte{0b01111111, 0b01111111, 0b10101010, 0b01111111, 0b00000011, 0b11111111}, e.Labels())
}
