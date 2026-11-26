package eclipse

// OR computes logical OR of labels.
func OR(f1, f2 []byte) []byte {
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
				e.SetByte(f1I, f1L|f2L)
				f1L, f1I, f1Ok = f1It.Next()
				f2L, f2I, f2Ok = f2It.Next()
			case f1I < f2I:
				e.SetByte(f1I, f1L)
				f1L, f1I, f1Ok = f1It.Next()
			default:
				e.SetByte(f2I, f2L)
				f2L, f2I, f2Ok = f2It.Next()
			}
		// FIXME (wojciech): The following two cases might be optimized by simply copying the rest of labels.
		case f1Ok:
			e.SetByte(f1I, f1L)
			f1L, f1I, f1Ok = f1It.Next()
		default:
			e.SetByte(f2I, f2L)
			f2L, f2I, f2Ok = f2It.Next()
		}
	}

	return e.Labels()
}

// AND computes logical AND of labels.
func AND(f1, f2 []byte) []byte {
	e := newEncoder()
	f1It := newIterator(f1)
	f2It := newIterator(f2)

	f2L, f2I, f2Ok := f2It.Next()
	f1L, f1I, f1Ok := f1It.Forward(f2I)

	for f1Ok && f2Ok {
		switch {
		case f1I == f2I:
			if v := f1L & f2L; v > oneByteNumFlag {
				e.SetByte(f1I, f1L&f2L)
			}
			f1L, f1I, f1Ok = f1It.Next()
			f2L, f2I, f2Ok = f2It.Forward(f1I)
		case f1I < f2I:
			f1L, f1I, f1Ok = f1It.Forward(f2I)
		default:
			f2L, f2I, f2Ok = f2It.Forward(f1I)
		}
	}

	return e.Labels()
}
