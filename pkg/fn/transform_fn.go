package fn

import (
	"math/rand"
	"sort"

	"github.com/Yangruipis/go-functional/pkg/iter"
)

func NewEntry[K, V any](k K, v V) iter.Entry[K, V] {
	return iter.Entry[K, V]{
		K: k,
		V: v,
	}
}

func Map[K, V any, K1, V1 any](i iter.Iterator[K, V], f func(k K, v V) (K1, V1)) iter.Iterator[K1, V1] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[K1, V1], e iter.Entry[K, V]) {
		k, v := f(e.K, e.V)
		c <- NewEntry(k, v)
	})
}

func Filter[K, V any](i iter.Iterator[K, V], f func(k K, v V) bool) iter.Iterator[K, V] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[K, V], e iter.Entry[K, V]) {
		keep := f(e.K, e.V)
		if keep {
			c <- e
		}
	})
}

func Flatten[K, V any](i iter.Iterator[K, []V]) iter.Iterator[K, V] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[K, V], e iter.Entry[K, []V]) {
		for _, v := range e.V {
			c <- NewEntry(e.K, v)
		}
	})
}

// XXX: not lazy
func GroupByKey[K iter.Comparable, V any](i iter.Iterator[K, V]) iter.Iterator[K, []V] {

	keys := make([]K, 0, 16)
	m := make(map[K][]V)

	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if _, ok := m[v.K]; !ok {
			m[v.K] = make([]V, 0, 32)
			keys = append(keys, v.K)
		}
		m[v.K] = append(m[v.K], v.V)
	}

	return iter.NewMapIteratorWithKeys(keys, m)
}

func GroupBy[K iter.Comparable, V any](i iter.Iterator[K, V], f func(k K, v V) K) iter.Iterator[K, []V] {
	return GroupByKey(Map(i, func(k K, v V) (K, V) {
		return f(k, v), v
	}))
}

func FlatMap[K, V any](i iter.Iterator[K, []V], f func(k K, v []V) (K, []V)) iter.Iterator[K, V] {
	return Flatten(Map(i, func(k K, v []V) (K, []V) {
		return f(k, v)
	}))

}

func SliceIter[V any](arr []V) iter.Iterator[int, V] {
	return iter.NewSliceIterator(arr)
}

func MapIter[K iter.Comparable, V any](m map[K]V) iter.Iterator[K, V] {
	return iter.NewMapIterator(m)
}

func Range(start, end, step int) iter.Iterator[int, int] {
	if step == 0 || start == end || (start < end && step < 0) || (start > end && step > 0) {
		return iter.NewSliceIterator([]int{})
	}

	reverseFlag := false
	if start > end {
		reverseFlag = true
	}

	c := make(chan iter.Entry[int, int], 1)
	go func() {
		i, idx := start, 0
		for {
			if (reverseFlag && i > end) || (!reverseFlag && i < end) {
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

func Repeat[V any](t V, num int) iter.Iterator[int, V] {
	if num <= 0 {
		return iter.NewSliceIterator([]V{})
	}
	c := make(chan iter.Entry[int, V], 1)

	go func() {
		for i := 0; i < num; i++ {
			c <- NewEntry(i, t)

		}
		close(c)

	}()
	return iter.NewChanIterator(c)
}

func ReduceByKey[K iter.Comparable, V any](i iter.Iterator[K, V], f func(a, b V) V) iter.Iterator[K, V] {
	return Map(GroupByKey(i), func(k K, v []V) (K, V) {
		return k, Reduce(SliceIter(v), f)
	})
}

func CountByKey[K iter.Comparable, V any](i iter.Iterator[K, V]) iter.Iterator[K, int] {
	return Map(GroupByKey(i), func(k K, v []V) (K, int) {
		return k, len(v)
	})
}

func Invert[K, V any](i iter.Iterator[K, V]) iter.Iterator[V, K] {
	return iter.NewChanIteratorF(i, func(c chan iter.Entry[V, K], e iter.Entry[K, V]) {
		c <- NewEntry(e.V, e.K)
	})
}

// not lazy
func Reverse[K, V any](i iter.Iterator[K, V]) iter.Iterator[K, V] {
	s := Entries(i)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return iter.NewEntryIterator(s)
}

func GroupByVal[K any, V iter.Comparable](i iter.Iterator[K, V]) iter.Iterator[V, []K] {
	return GroupByKey(Invert(i))
}

func CountByVal[K any, V iter.Comparable](i iter.Iterator[K, V]) iter.Iterator[V, int] {
	return CountByKey(Invert(i))
}

func UnionBy[K, V any, K1 iter.Comparable](i1, i2 iter.Iterator[K, V], f func(k K, v V) K1) iter.Iterator[K, V] {
	m := make(map[K1]struct{})
	c := make(chan iter.Entry[K, V], 1)

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

func IntersectBy[K, V any, K1 iter.Comparable](i1, i2 iter.Iterator[K, V], f func(k K, v V) K1) iter.Iterator[K, V] {
	m := make(map[K1]struct{})
	c := make(chan iter.Entry[K, V], 1)

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
			if _, ok := m[key]; ok {
				c <- v
			}
		}
		close(c)
	}()
	return iter.NewChanIterator(c)
}

func SubstractBy[K, V any, K1 iter.Comparable](i1, i2 iter.Iterator[K, V], f func(k K, v V) K1) iter.Iterator[K, V] {
	m := make(map[K1]struct{})
	c := make(chan iter.Entry[K, V], 1)

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

func DistinctBy[K, V any, K1 iter.Comparable](i iter.Iterator[K, V], f func(k K, v V) K1) iter.Iterator[K, V] {

	m := make(map[K1]struct{})
	c := make(chan iter.Entry[K, V], 1)

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

func Union[K any, V iter.Comparable](i1, i2 iter.Iterator[K, V]) iter.Iterator[K, V] {
	return UnionBy(i1, i2, func(k K, v V) V {
		return v
	})
}

func Intersect[K any, V iter.Comparable](i1, i2 iter.Iterator[K, V]) iter.Iterator[K, V] {
	return IntersectBy(i1, i2, func(k K, v V) V {
		return v
	})
}

func Subtract[K any, V iter.Comparable](i1, i2 iter.Iterator[K, V]) iter.Iterator[K, V] {
	return SubstractBy(i1, i2, func(k K, v V) V {
		return v
	})
}

func Distinct[K any, V iter.Comparable](i iter.Iterator[K, V]) iter.Iterator[K, V] {
	return DistinctBy(i, func(k K, v V) V {
		return v
	})
}

func Cartesian[K any, V, V1 any](i1, i2 iter.Iterator[K, V], f func(v1, v2 V) V1) iter.Iterator[int, V1] {
	c := make(chan iter.Entry[int, V1], 1)
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

func Chunk[K iter.Comparable, V any](i iter.Iterator[K, V], size int) iter.Iterator[int, []V] {
	if size <= 0 {
		panic("chunk size must be positiive")
	}
	c := make(chan iter.Entry[int, []V], 1)

	go func() {
		idx := 0
		chunks := []V{}
		for {
			v, flag := i.Next()
			if (flag == iter.FlagStop && len(chunks) > 0) || len(chunks) == size {
				c <- NewEntry(idx, chunks)
				chunks = []V{}
				idx++
			}
			if flag == iter.FlagStop {
				close(c)
				return
			}
			chunks = append(chunks, v.V)

		}
	}()
	return iter.NewChanIterator(c)
}

func Sort[K, V any](i iter.Iterator[K, V], lessFn func(v1, v2 V) bool) iter.Iterator[K, V] {
	s := Entries(i)
	sort.Slice(s, func(i, j int) bool {
		return lessFn(s[i].V, s[j].V)
	})
	return iter.NewEntryIterator(s)
}

func Aggregate[K iter.Comparable, V, V1 any](i iter.Iterator[K, []V], f func(vv []V) V1) iter.Iterator[K, V1] {
	return Map(i, func(k K, v []V) (K, V1) {
		return k, f(v)
	})
}

func Zip[V, V1 any](i1 iter.Iterator[int, V], i2 iter.Iterator[int, V1]) iter.Iterator[int, iter.Entry[V, V1]] {
	c := make(chan iter.Entry[int, iter.Entry[V, V1]], 1)

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

func Shuffle[K, V any](i iter.Iterator[K, V]) iter.Iterator[K, V] {
	s := Entries(i)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return iter.NewEntryIterator(s)
}

// without replacement, which means element0 will be sampled only once
func Sample[K, V any](i iter.Iterator[K, V], size float32) iter.Iterator[K, V] {
	if size <= 0 {
		panic("size must be positive")
	}
	s := Entries(i)
	intSize := int(size)
	if intSize > len(s) {
		intSize = len(s)
	} else if size < 1 {
		intSize = int(size * float32(len(s)))
	}
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return iter.NewEntryIterator(s[:intSize])
}

// with replacement, which means element0 will be sampled any times
func Choices[K, V any](i iter.Iterator[K, V], size float32) iter.Iterator[K, V] {
	if size <= 0 {
		panic("size must be positive")
	}
	s := Entries(i)
	intSize := int(size)
	if size < 1 {
		intSize = int(size * float32(len(s)))
	}

	c := make(chan iter.Entry[K, V], 1)
	go func() {
		for i := 0; i < intSize; i++ {
			tmp := s[rand.Intn(len(s))]
			c <- NewEntry(tmp.K, tmp.V)
		}
		close(c)
	}()
	return iter.NewChanIterator(c)
}

func Head[K, V any](i iter.Iterator[K, V], n int) iter.Iterator[K, V] {
	c := make(chan iter.Entry[K, V], 1)

	go func() {
		idx := 0
		for {
			v, flag := i.Next()
			if flag == iter.FlagStop || idx >= n {
				close(c)
				return
			}
			c <- NewEntry(v.K, v.V)
			idx++
		}
	}()
	return iter.NewChanIterator(c)
}

func Tail[K, V any](i iter.Iterator[K, V], n int) iter.Iterator[K, V] {
	return Head(Reverse(i), n)
}

func Cache[K, V any](i iter.Iterator[K, V]) iter.Iterator[K, V] {
	return iter.NewCachedIterator(Entries(i))
}
