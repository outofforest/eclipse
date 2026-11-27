package eclipse

type iterator struct {
	labels    []uint64
	realIndex uint64
}

func newIterator(labels []uint64) *iterator {
	return &iterator{
		labels: labels,
	}
}

func (it *iterator) Next() (uint64, uint64, bool) {
	return it.Forward(0)
}

func (it *iterator) Forward(index uint64) (uint64, uint64, bool) {
	for i, l := range it.labels {
		if l&oneByteNumFlag == 0 {
			it.realIndex += l
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
