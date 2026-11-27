package eclipse

import (
	"math"
)

// LabelSource is a source of labels.
type LabelSource func() []uint64

// Direct creates label source from label slice.
func Direct(labels []uint64) LabelSource {
	return func() []uint64 {
		return labels
	}
}

// OR computes logical OR of labels.
func OR(fSources ...LabelSource) LabelSource {
	var cache []uint64
	return func() []uint64 {
		if cache != nil {
			return cache
		}

		its := make([]*iterator, 0, len(fSources))
		for _, fSource := range fSources {
			f := fSource()
			if len(f) > 0 {
				cache = f
				its = append(its, newIterator(f))
			}
		}
		switch len(its) {
		case 0:
			cache = []uint64{}
			return cache
		case 1:
			return cache
		}

		e := newEncoder()

		incrementIts := make([]*iterator, 0, len(its))
		for {
			incrementIts = incrementIts[:0]
			var ok bool
			var minFI uint64 = math.MaxUint64
			var l uint64
			for i, it := range its {
				if it == nil {
					continue
				}

				fL, fI, fOk := it.Next()
				if !fOk {
					its[i] = nil
					continue
				}

				ok = true
				switch {
				case fI < minFI:
					minFI = fI
					l = fL
					incrementIts = incrementIts[:0]
					incrementIts = append(incrementIts, it)
				case fI == minFI:
					l |= fL
					incrementIts = append(incrementIts, it)
				}
			}
			if !ok {
				break
			}

			e.SetChunk(minFI, l)
			for _, it := range incrementIts {
				it.Next()
			}
		}

		cache = e.Labels()
		return cache
	}
}

// AND computes logical AND of labels.
func AND(f1Source, f2Source LabelSource) LabelSource {
	var cache []uint64
	return func() []uint64 {
		if cache != nil {
			return cache
		}

		f1 := f1Source()
		if len(f1) == 0 {
			cache = f1
			return cache
		}

		f2 := f2Source()
		if len(f2) == 0 {
			cache = f2
			return cache
		}

		e := newEncoder()
		f1It := newIterator(f1)
		f2It := newIterator(f2)

		f2L, f2I, f2Ok := f2It.Next()
		f1L, f1I, f1Ok := f1It.Forward(f2I)

		for f1Ok && f2Ok {
			switch {
			case f1I == f2I:
				if v := f1L & f2L; v > oneByteNumFlag {
					e.SetChunk(f1I, f1L&f2L)
				}
				f1L, f1I, f1Ok = f1It.Next()
				f2L, f2I, f2Ok = f2It.Forward(f1I)
			case f1I < f2I:
				f1L, f1I, f1Ok = f1It.Forward(f2I)
			default:
				f2L, f2I, f2Ok = f2It.Forward(f1I)
			}
		}

		cache = e.Labels()
		return cache
	}
}
