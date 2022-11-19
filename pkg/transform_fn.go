package fun

import (
	"math/rand"
	"sort"

	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

func NewEntry[T1, T2 any](k T1, v T2) iter.Entry[T1, T2] {
	return iter.Entry[T1, T2]{
		K: k,
		V: v,
	}
}

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

func Invert[T1, T2 any](i iter.Iterator[T1, T2]) iter.Iterator[T2, T1] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[T2, T1], e iter.Entry[T1, T2]) {
		c <- NewEntry(e.V, e.K)
	})
}

// not lazy
func Reverse[T2 any](i iter.Iterator[int, T2]) iter.Iterator[int, T2] {
	s := ToSlice(i)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return iter.NewSliceIterator[T2](s)
}

func GroupByVal[T1 any, T2 iter.Hashable](i iter.Iterator[T1, T2]) iter.Iterator[T2, []T1] {
	return GroupByKey(Invert(i))
}

func CountByVal[T1 any, T2 iter.Hashable](i iter.Iterator[T1, T2]) iter.Iterator[T2, int] {
	return CountByKey(Invert(i))
}

func UnionBy[T1, T2 any, T3 iter.Hashable](i1, i2 iter.Iterator[T1, T2], f func(k T1, v T2) T3) iter.Iterator[T1, T2] {
	m := make(map[T3]struct{})
	c := make(chan iter.Entry[T1, T2], 1)

	go func() {
		for {
			v, flag := i1.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			if _, ok := m[key]; !ok {
				c <- v
				m[key] = struct{}{}
			}
		}
		for {
			v, flag := i2.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			if _, ok := m[key]; !ok {
				c <- v
				m[key] = struct{}{}
			}
		}
		close(c)
	}()
	return iter.NewChanIterator(c)
}

func IntersectBy[T1, T2 any, T3 iter.Hashable](i1, i2 iter.Iterator[T1, T2], f func(k T1, v T2) T3) iter.Iterator[T1, T2] {
	m := make(map[T3]struct{})
	c := make(chan iter.Entry[T1, T2], 1)

	go func() {
		for {
			v, flag := i1.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			m[key] = struct{}{}
		}
		for {
			v, flag := i2.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			if _, ok := m[key]; ok {
				c <- v
			}
		}
		close(c)
	}()
	return iter.NewChanIterator(c)
}

func SubstractBy[T1, T2 any, T3 iter.Hashable](i1, i2 iter.Iterator[T1, T2], f func(k T1, v T2) T3) iter.Iterator[T1, T2] {
	m := make(map[T3]struct{})
	c := make(chan iter.Entry[T1, T2], 1)

	go func() {
		for {
			v, flag := i2.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			m[key] = struct{}{}
		}
		for {
			v, flag := i1.Next()
			if flag == iter.FlagStop {
				break
			}
			key := f(v.K, v.V)
			if _, ok := m[key]; !ok {
				c <- v
			}
		}
		close(c)

	}()
	return iter.NewChanIterator(c)
}

func DistinctBy[T1, T2 any, T3 iter.Hashable](i iter.Iterator[T1, T2], f func(k T1, v T2) T3) iter.Iterator[T1, T2] {

	m := make(map[T3]struct{})
	c := make(chan iter.Entry[T1, T2], 1)

	go func() {
		for {
			v, flag := i.Next()
			if flag == iter.FlagStop {
				close(c)
				return
			}
			key := f(v.K, v.V)
			if _, ok := m[key]; !ok {
				c <- v
				m[key] = struct{}{}
			}
		}

	}()
	return iter.NewChanIterator(c)

}

func Union[T1 any, T2 iter.Hashable](i1, i2 iter.Iterator[T1, T2]) iter.Iterator[T1, T2] {
	return UnionBy(i1, i2, func(k T1, v T2) T2 {
		return v
	})
}

func Intersect[T1 any, T2 iter.Hashable](i1, i2 iter.Iterator[T1, T2]) iter.Iterator[T1, T2] {
	return IntersectBy(i1, i2, func(k T1, v T2) T2 {
		return v
	})
}

func Subtract[T1 any, T2 iter.Hashable](i1, i2 iter.Iterator[T1, T2]) iter.Iterator[T1, T2] {
	return SubstractBy(i1, i2, func(k T1, v T2) T2 {
		return v
	})
}

func Distinct[T1 any, T2 iter.Hashable](i iter.Iterator[T1, T2]) iter.Iterator[T1, T2] {
	return DistinctBy(i, func(k T1, v T2) T2 {
		return v
	})
}

func Cartesian[T1 int, T2, O2 any](i1, i2 iter.Iterator[T1, T2], f func(v1, v2 T2) O2) iter.Iterator[int, O2] {
	c := make(chan iter.Entry[int, O2], 1)
	vv2 := ToSlice(i2)

	go func() {
		idx := 0
		for {
			v1, flag := i1.Next()
			if flag == iter.FlagStop {
				close(c)
				return
			}
			for _, v2 := range vv2 {
				v := f(v1.V, v2)
				c <- NewEntry(idx, v)
				idx++
			}
		}
	}()
	return iter.NewChanIterator(c)
}

func Chunk[T1 iter.Hashable, T2 any](i iter.Iterator[T1, T2], size int) iter.Iterator[int, []T2] {
	if size <= 0 {
		panic("chunk size must be positivei")
	}
	c := make(chan iter.Entry[int, []T2], 1)

	go func() {
		idx := 0
		chunks := []T2{}
		for {
			v, flag := i.Next()
			if (flag == iter.FlagStop && len(chunks) > 0) || len(chunks) == size {
				c <- NewEntry(idx/size, chunks)
				chunks = []T2{}
			}
			if flag == iter.FlagStop {
				close(c)
				return
			}
			chunks = append(chunks, v.V)
			idx++

		}
	}()
	return iter.NewChanIterator(c)
}

func Sort[T1, T2 any](i iter.Iterator[T1, T2], lessFn func(v1, v2 T2) bool) iter.Iterator[int, T2] {
	s := ToSlice(i)
	sort.Slice(s, func(i, j int) bool {
		return lessFn(s[i], s[j])
	})
	return iter.NewSliceIterator(s)
}

func Aggregate[T1 iter.Hashable, T2, O any](i iter.Iterator[T1, T2], f func([]T2) O) iter.Iterator[T1, O] {
	return Map(GroupByKey(i), func(k T1, v []T2) (T1, O) {
		return k, f(v)
	})
}

func Zip[T, O any](i1 iter.Iterator[int, T], i2 iter.Iterator[int, O]) iter.Iterator[int, iter.Entry[T, O]] {
	c := make(chan iter.Entry[int, iter.Entry[T, O]], 1)

	go func() {
		idx := 0
		for {
			v1, flag1 := i1.Next()
			v2, flag2 := i2.Next()
			if flag1 == iter.FlagStop || flag2 == iter.FlagStop {
				close(c)
				return
			}

			c <- NewEntry(idx, NewEntry(v1.V, v2.V))
			idx++
		}
	}()
	return iter.NewChanIterator(c)
}

func Shuffle[T2 any](i iter.Iterator[int, T2]) iter.Iterator[int, T2] {
	s := ToSlice(i)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return iter.NewSliceIterator(s)
}

func Choice[T2 any](i iter.Iterator[int, T2], size float32) iter.Iterator[int, T2] {
	if size <= 0 {
		panic("size must be positive")
	}
	s := ToSlice(i)
	intSize := int(size)
	if intSize > len(s) {
		intSize = len(s)
	} else if size < 1 {
		intSize = int(size * float32(len(s)))
	}
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return iter.NewSliceIterator(s[:intSize])
}

func Sample[T2 any](i iter.Iterator[int, T2], size float32) iter.Iterator[int, T2] {
	if size <= 0 {
		panic("size must be positive")
	}
	s := ToSlice(i)
	intSize := int(size)
	if size < 1 {
		intSize = int(size * float32(len(s)))
	}

	c := make(chan iter.Entry[int, T2], 1)
	go func() {
		for i := 0; i < intSize; i++ {
			c <- NewEntry(i, s[rand.Intn(len(s))])
		}
		idx := 0
	}()
	return iter.NewChanIterator(c)
}
