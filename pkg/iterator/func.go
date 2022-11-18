package iter

type FuncIterator[T1, T2 any] struct {
	Iter func() (Entry[T1, T2], Flag)
}

func NewFuncIterator[T1, T2 any](f func() (Entry[T1, T2], Flag)) *FuncIterator[T1, T2] {
	return &FuncIterator[T1, T2]{
		Iter: f,
	}
}

func (i *FuncIterator[T1, T2]) Next() (v Entry[T1, T2], flag Flag) {
	v, flag = i.Iter()
	if flag == FlagStop {
		return
	}
	if flag == FlagSkip {
		return i.Next()
	}
	return
}
