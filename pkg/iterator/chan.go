package iter

type ChanIterator[T1, T2 any] struct {
	Iter <-chan Entry[T1, T2]
}

func NewChanIterator[T1, T2 any](c <-chan Entry[T1, T2]) *ChanIterator[T1, T2] {
	return &ChanIterator[T1, T2]{
		Iter: c,
	}
}

func NewChanIteratorF[T1, T2, O1, O2 any](i Iterator[T1, T2], f func(c chan Entry[O1, O2], e Entry[T1, T2])) *ChanIterator[O1, O2] {
	c := make(chan Entry[O1, O2], 1)
	go func() {
		for {
			v, flag := i.Next()
			if flag == FlagStop {
				close(c)
				return
			}
			f(c, v)
		}
	}()
	return NewChanIterator(c)
}

func (i *ChanIterator[T1, T2]) Next() (v Entry[T1, T2], flag Flag) {
	v, ok := (<-i.Iter)
	if !ok {
		flag = FlagStop
		return
	}
	return
}
