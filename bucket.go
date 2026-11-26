package eclipse

// ValueRangeFunc is the function returning range of value.
type ValueRangeFunc[T any, U Bucketable] func(o T) *Range[U]

// Range defines range of value.
type Range[T Bucketable] struct {
	From T
	To   T
}

// Covers returns true if r2 is fully in the range r.
func (r Range[T]) Covers(value T) bool {
	return value >= r.From && value <= r.To
}

// FromTo produces range based on from and to boundaries.
func FromTo[U Bucketable](from, to U) Range[U] {
	return Range[U]{From: from, To: to}
}

// Point produces range containing single exact value.
func Point[U Bucketable](v U) Range[U] {
	return FromTo[U](v, v)
}

// Bucketable represents types which might be used in bucket.
type Bucketable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

var _ Filter[struct{}] = &bucketFilter[struct{}, int]{}

type bucketFilter[T any, U Bucketable] struct {
	minMax         Range[U]
	valueRangeFunc ValueRangeFunc[T, U]
	bucketFactor   float64
	maxBucketIndex uint64
	encoders       []*encoder
	index          uint64
}

// NewBucketFilter creates new bucket filter.
func NewBucketFilter[T any, U Bucketable](
	numOfBuckets uint64,
	minMax Range[U],
	valueRangeFunc ValueRangeFunc[T, U],
) Filter[T] {
	bf := &bucketFilter[T, U]{
		minMax:         minMax,
		valueRangeFunc: valueRangeFunc,
		bucketFactor:   float64(numOfBuckets) / float64(minMax.To-minMax.From),
		maxBucketIndex: numOfBuckets - 1,
		encoders:       make([]*encoder, numOfBuckets),
	}

	for i := range numOfBuckets {
		bf.encoders[i] = newEncoder()
	}

	return bf
}

func (b *bucketFilter[T, U]) Add(o T) {
	minBucket, maxBucket, ok := b.buckets(o)
	if ok {
		for i := minBucket; i <= maxBucket; i++ {
			b.encoders[i].SetBit(b.index)
		}
	}
	b.index++
}

func (b *bucketFilter[T, U]) Match(o T) []byte {
	minBucket, maxBucket, ok := b.buckets(o)
	if !ok {
		return nil
	}

	labels := b.encoders[minBucket].Labels()
	for i := minBucket + 1; i <= maxBucket; i++ {
		labels = OR(labels, b.encoders[i].Labels())
	}
	return labels
}

func (b *bucketFilter[T, U]) buckets(o T) (uint64, uint64, bool) {
	valueRange := b.valueRangeFunc(o)
	if valueRange == nil {
		return 0, 0, false
	}

	from := max(valueRange.From, b.minMax.From)
	to := max(valueRange.To, b.minMax.From)

	minBucket := min(uint64(float64(from-b.minMax.From)*b.bucketFactor), b.maxBucketIndex)
	maxBucket := min(uint64(float64(to-b.minMax.From)*b.bucketFactor), b.maxBucketIndex)

	return minBucket, maxBucket, true
}
