package eclipse

import "iter"

// Filter provides bit filter for specific criterion.
type Filter[T any] interface {
	Add(o T)
	Match(o T) []byte
}

// Matcher matches records against filters.
type Matcher[T any] struct {
	filters []Filter[T]
}

// NewMatcher creates new matcher.
func NewMatcher[T any](filters ...Filter[T]) *Matcher[T] {
	return &Matcher[T]{
		filters: filters,
	}
}

// Add adds record to filters.
func (m *Matcher[T]) Add(o T) {
	for _, f := range m.filters {
		f.Add(o)
	}
}

// Match finds records matching provided criteria.
func (m *Matcher[T]) Match(o T) iter.Seq[int] {
	var labels []byte
	for _, f := range m.filters {
		l := f.Match(o)
		if len(l) == 0 {
			if l == nil {
				continue
			}
			return func(yield func(int) bool) {}
		}
		if labels == nil {
			labels = l
		} else {
			labels = AND(labels, l)
		}
	}
	if labels == nil {
		return nil
	}

	it := newIterator(labels)
	return func(yield func(int) bool) {
		for {
			l, i, ok := it.Next()
			if !ok {
				return
			}
			bitIndex := int(i * numOfBits)
			for j := range numOfBits {
				mask := byte(1) << j
				if mask > l {
					break
				}
				if l&mask != 0 && !yield(bitIndex+j) {
					return
				}
			}
		}
	}
}
