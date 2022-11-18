package fun

import iter "github.com/Yangruipis/go-functional/pkg/iterator"

func Map[T1, T2 any, O1, O2 any](i iter.Iterator[T1, T2], f func(k T1, v T2) (O1, O2)) iter.Iterator[O1, O2] {
	ff := func() (v iter.Entry[O1, O2], flag iter.Flag) {
		vSrc, flag := i.Next()
		if flag == iter.FlagStop {
			return
		}
		vDstK, vDstV := f(vSrc.K, vSrc.V)
		return iter.Entry[O1, O2]{
			K: vDstK,
			V: vDstV,
		}, iter.FlagOK
	}
	return iter.NewFuncIterator(ff)
}

func Filter[T1, T2 any](i iter.Iterator[T1, T2], f func(k T1, v T2) bool) iter.Iterator[T1, T2] {
	ff := func() (v iter.Entry[T1, T2], flag iter.Flag) {
		vSrc, flag := i.Next()
		if flag == iter.FlagStop {
			return
		}
		keep := f(vSrc.K, vSrc.V)
		if !keep {
			return vSrc, iter.FlagSkip
		}
		return vSrc, iter.FlagOK

	}
	return iter.NewFuncIterator(ff)
}

func Flatten[T1, T2 any](i iter.Iterator[T1, []T2]) iter.Iterator[T1, T2] {

	ff := func() func() (v iter.Entry[T1, T2], flag iter.Flag) {

		vv := iter.Entry[T1, []T2]{}
		idx := 0

		return func() (v iter.Entry[T1, T2], flag iter.Flag) {
			if idx >= len(vv.V) {
				vv, flag = i.Next()
				idx = 0
				if flag == iter.FlagStop {
					return
				}
			}

			rtn := iter.Entry[T1, T2]{
				K: vv.K,
				V: vv.V[idx],
			}
			idx++
			return rtn, iter.FlagOK
		}
	}
	return iter.NewFuncIterator(ff())
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
	i := start
	idx := 0
	f := func() (k, v int, flag iter.Flag) {
		if i < end {
			v = i
			k = idx
			i += step
			idx++
		} else {
			flag = iter.FlagStop
		}

		return
	}
	return NewGenerator(f)
}

func Repeat[T any](t T, num int) iter.Iterator[int, T] {
	idx := 0
	f := func() (k int, v T, flag iter.Flag) {
		if idx < num {
			v = t
			idx++
		} else {
			flag = iter.FlagStop
		}

		return
	}
	return NewGenerator(f)
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
