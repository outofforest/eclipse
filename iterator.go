package eclipse

const oneByteNumFlag = 128

type iterator struct {
	labels    []byte
	realIndex uint64
}

func newIterator(labels []byte) *iterator {
	return &iterator{
		labels: labels,
	}
}

func (it *iterator) Next() (byte, uint64, bool) {
	return it.Forward(0)
}

func (it *iterator) Forward(index uint64) (byte, uint64, bool) {
	for i, l := range it.labels {
		if l&oneByteNumFlag == 0 {
			it.realIndex += uint64(l)
			continue
		}

		it.realIndex++
		if it.realIndex > index {
			it.labels = it.labels[i+1:]
			return l, it.realIndex - 1, true
		}
	}
	return 0x00, 0, false
}
