package fun

import iter "github.com/Yangruipis/go-functional/pkg/iterator"

func Map[T1, T2 any, O1, O2 any](i iter.Iterator[T1, T2], f func(k T1, v T2) (O1, O2)) iter.Iterator[O1, O2] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[O1, O2], e iter.Entry[T1, T2]) {
		k, v := f(e.K, e.V)
		c <- NewEntry(k, v)
	})
}

func Filter[T1, T2 any](i iter.Iterator[T1, T2], f func(k T1, v T2) bool) iter.Iterator[T1, T2] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[T1, T2], e iter.Entry[T1, T2]) {
		keep := f(e.K, e.V)
		if keep {
			c <- e
		}
	})
}

func Flatten[T1, T2 any](i iter.Iterator[T1, []T2]) iter.Iterator[T1, T2] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[T1, T2], e iter.Entry[T1, []T2]) {
		for _, v := range e.V {
			c <- NewEntry(e.K, v)
		}
	})
}

// XXX: not lazy
func GroupByKey[T1 iter.Hashable, T2 any](i iter.Iterator[T1, T2]) iter.Iterator[T1, []T2] {

	m := make(map[T1][]T2)

	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if _, ok := m[v.K]; !ok {
			m[v.K] = make([]T2, 0, 32)
		}
		m[v.K] = append(m[v.K], v.V)
	}

	return iter.NewMapIterator(m)
}

func GroupBy[T1 iter.Hashable, T2 any](i iter.Iterator[T1, T2], f func(k T1) T1) iter.Iterator[T1, []T2] {
	return GroupByKey(Map(i, func(k T1, v T2) (T1, T2) {
		return f(k), v
	}))
}

func FlatMap[T1, T2 any](i iter.Iterator[T1, []T2], f func(k T1, v []T2) (T1, []T2)) iter.Iterator[T1, T2] {
	return Flatten(Map(i, func(k T1, v []T2) (T1, []T2) {
		return f(k, v)
	}))

}

func Range(start, end, step int) iter.Iterator[int, int] {
	c := make(chan iter.Entry[int, int], 1)

	go func() {
		i, idx := start, 0
		for {
			if i < end {
				c <- NewEntry(idx, i)
				i += step
				idx++
			} else {
				close(c)
				return
			}
		}

	}()
	return iter.NewChanIterator(c)
}

func Repeat[T any](t T, num int) iter.Iterator[int, T] {
	c := make(chan iter.Entry[int, T], 1)

	go func() {
		for i := 0; i < num; i++ {
			c <- NewEntry(i, t)

		}
		close(c)

	}()
	return iter.NewChanIterator(c)
}

func ReduceByKey[T1 iter.Hashable, T2 any](i iter.Iterator[T1, T2], f func(a, b T2) T2) iter.Iterator[T1, T2] {
	return Map(GroupByKey(i), func(k T1, v []T2) (T1, T2) {
		return k, Reduce(NewSlice(v), f)
	})
}

func CountByKey[T1 iter.Hashable, T2 any](i iter.Iterator[T1, T2]) iter.Iterator[T1, int] {
	return Map(GroupByKey(i), func(k T1, v []T2) (T1, int) {
		return k, len(v)
	})
}

func NewEntry[T1, T2 any](k T1, v T2) iter.Entry[T1, T2] {
	return iter.Entry[T1, T2]{
		K: k,
		V: v,
	}
}
