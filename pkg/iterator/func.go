package iter

type FuncIterator[K, V any] struct {
	Iter func() (Entry[K, V], Flag)
}

func NewFuncIterator[K, V any](f func() (Entry[K, V], Flag)) *FuncIterator[K, V] {
	return &FuncIterator[K, V]{
		Iter: f,
	}
}

func (i *FuncIterator[K, V]) Next() (v Entry[K, V], flag Flag) {
	v, flag = i.Iter()
	if flag == FlagStop {
		return
	}
	if flag == FlagSkip {
		return i.Next()
	}
	return
}
