package eclipse

// ValueDictionaryFunc returns dictionary from object.
type ValueDictionaryFunc[T, U any] func(o T) []U

var _ Filter[struct{}] = &dictionaryFilter[struct{}, int]{}

type dictionaryFilter[T any, U comparable] struct {
	dictFunc  ValueDictionaryFunc[T, U]
	logicFunc func(f1, f2 []byte) []byte
	encoders  map[U]*encoder
	index     uint64
}

// NewDictionaryFilter returns filter creating map for each item in dictionary.
func NewDictionaryFilter[T any, U comparable](
	dictFunc ValueDictionaryFunc[T, U],
	logicFunc func(f1, f2 []byte) []byte,
) Filter[T] {
	return &dictionaryFilter[T, U]{
		dictFunc:  dictFunc,
		logicFunc: logicFunc,
		encoders:  map[U]*encoder{},
	}
}

func (d *dictionaryFilter[T, U]) Add(o T) {
	values := d.dictFunc(o)
	for _, v := range values {
		e, exists := d.encoders[v]
		if !exists {
			e = newEncoder()
			d.encoders[v] = e
		}
		e.SetBit(d.index)
	}
	d.index++
}

func (d *dictionaryFilter[T, U]) Match(o T) []byte {
	values := d.dictFunc(o)
	if len(values) == 0 {
		if values == nil {
			return nil
		}
		return []byte{}
	}
	var labels []byte
	for _, v := range values {
		e := d.encoders[v]
		if e == nil {
			continue
		}
		if labels == nil {
			labels = e.Labels()
		} else {
			labels = d.logicFunc(labels, e.Labels())
		}
	}

	return labels
}
