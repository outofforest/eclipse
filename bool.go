package eclipse

// ValueBoolFunc returns bool value from object.
type ValueBoolFunc[T any] func(o T) bool

var _ Filter[struct{}] = &boolFilter[struct{}]{}

type boolFilter[T any] struct {
	boolFunc ValueBoolFunc[T]
	encoder  *encoder
	index    uint64
}

// NewBoolFilter creates new boolean filter.
func NewBoolFilter[T any](boolFunc ValueBoolFunc[T]) Filter[T] {
	return &boolFilter[T]{
		boolFunc: boolFunc,
		encoder:  newEncoder(),
	}
}

func (b *boolFilter[T]) Add(o T) {
	if b.boolFunc(o) {
		b.encoder.SetBit(b.index)
	}
	b.index++
}

func (b *boolFilter[T]) Match(o T) LabelSource {
	if b.boolFunc(o) {
		return Direct(b.encoder.Labels())
	}
	return nil
}
