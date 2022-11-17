package iter

type FuncIterator[T1, T2 any] struct {
	Iter func() (Entry[T1, T2], error)
}

func NewFuncIterator[T1, T2 any](f func() (Entry[T1, T2], error)) *FuncIterator[T1, T2] {
	return &FuncIterator[T1, T2]{
		Iter: f,
	}
}

func (i *FuncIterator[T1, T2]) Next() (v Entry[T1, T2], err error) {
	v, err = i.Iter()
	if err == StopIteration {
		return
	}
	if err == ContinueIteration {
		return i.Next()
	}
	return
}
