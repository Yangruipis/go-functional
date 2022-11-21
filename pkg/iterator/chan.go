package iter

type ChanIterator[K, V any] struct {
	Iter <-chan Entry[K, V]
}

func NewChanIterator[K, V any](c <-chan Entry[K, V]) *ChanIterator[K, V] {
	return &ChanIterator[K, V]{
		Iter: c,
	}
}

func NewChanIteratorF[K, V, K1, V1 any](i Iterator[K, V], f func(c chan Entry[K1, V1], e Entry[K, V])) *ChanIterator[K1, V1] {
	c := make(chan Entry[K1, V1], 1)
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

func (i *ChanIterator[K, V]) Next() (v Entry[K, V], flag Flag) {
	v, ok := (<-i.Iter)
	if !ok {
		flag = FlagStop
		return
	}
	return
}
