package eclipse

// LabelSource is a source of labels.
type LabelSource func() []uint64

// Direct creates label source from label slice.
func Direct(labels []uint64) LabelSource {
	return func() []uint64 {
		return labels
	}
}

// OR computes logical OR of labels.
func OR(f1Source, f2Source LabelSource) LabelSource {
	var cache []uint64
	return func() []uint64 {
		if cache != nil {
			return cache
		}

		f1 := f1Source()
		if len(f1) == 0 {
			cache = f2Source()
			return cache
		}

		f2 := f2Source()
		if len(f2) == 0 {
			cache = f1
			return cache
		}

		e := newEncoder()
		f1It := newIterator(f1)
		f2It := newIterator(f2)

		f1L, f1I, f1Ok := f1It.Next()
		f2L, f2I, f2Ok := f2It.Next()

		for f1Ok || f2Ok {
			switch {
			case f1Ok && f2Ok:
				switch {
				case f1I == f2I:
					e.SetChunk(f1I, f1L|f2L)
					f1L, f1I, f1Ok = f1It.Next()
					f2L, f2I, f2Ok = f2It.Next()
				case f1I < f2I:
					e.SetChunk(f1I, f1L)
					f1L, f1I, f1Ok = f1It.Next()
				default:
					e.SetChunk(f2I, f2L)
					f2L, f2I, f2Ok = f2It.Next()
				}
			// FIXME (wojciech): The following two cases might be optimized by simply copying the rest of labels.
			case f1Ok:
				e.SetChunk(f1I, f1L)
				f1L, f1I, f1Ok = f1It.Next()
			default:
				e.SetChunk(f2I, f2L)
				f2L, f2I, f2Ok = f2It.Next()
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
