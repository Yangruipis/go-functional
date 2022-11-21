package fn

import (
	"github.com/Yangruipis/go-functional/pkg/iter"
)

func ToSlice[K, V any](i iter.Iterator[K, V]) []V {
	res := make([]V, 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v.V)
	}
	return res
}

func ToMap[K iter.Comparable, V any](i iter.Iterator[K, V]) map[K]V {
	r := make(map[K]V)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		r[v.K] = v.V
	}
	return r
}

func ToSet[K any, V iter.Comparable](i iter.Iterator[K, V]) map[V]struct{} {
	r := make(map[V]struct{})
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		r[v.V] = struct{}{}
	}
	return r
}

func Values[K, V any](i iter.Iterator[K, V]) []V {
	return ToSlice(i)
}

func Keys[K, V any](i iter.Iterator[K, V]) []K {
	res := make([]K, 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v.K)
	}
	return res
}

func Entries[K, V any](i iter.Iterator[K, V]) []iter.Entry[K, V] {
	res := make([]iter.Entry[K, V], 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v)
	}
	return res
}

func ForEach[K, V any](i iter.Iterator[K, V], f func(k K, v V)) {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		f(v.K, v.V)
	}
}

func Reduce[K, V any](i iter.Iterator[K, V], f func(a, b V) V) V {
	var (
		res V
		idx int
	)

	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if idx == 0 {
			res = v.V
		} else {
			res = f(res, v.V)
		}
		idx += 1
	}
	return res
}

func Size[K, V any](i iter.Iterator[K, V]) int {
	return len(ToSlice(i))
}

func Any[K any, V bool](i iter.Iterator[K, V]) bool {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if v.V {
			return true
		}
	}
	return false
}

func All[K any, V bool](i iter.Iterator[K, V]) bool {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if !v.V {
			return false
		}
	}
	return true
}

func ExistsBy[K, V any](i iter.Iterator[K, V], eq func(vsrc V) bool) bool {
	return Any(Map(i, func(k K, v V) (K, bool) {
		return k, eq(v)
	}))
}

func Exists[K any, V iter.Comparable](i iter.Iterator[K, V], target V) bool {
	return ExistsBy(i, func(v V) bool {
		return v == target
	})
}

func Contains[K any, V iter.Comparable](i iter.Iterator[K, V], target V) bool {
	return Exists(i, target)
}

func Sum[K any, V iter.Comparable](i iter.Iterator[K, V]) (res V) {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res += v.V
	}
	return
}

func Avg[K any, V iter.Numeric | iter.Int | iter.Uint](i iter.Iterator[K, V]) (res V) {
	cnt := 0
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res += v.V
		cnt++
	}
	return res / V(cnt)
}

func CountBy[K, V any](i iter.Iterator[K, V], cmp func(v V) bool) int {
	return Sum(Map(i, func(k K, v V) (K, int) {
		if r := cmp(v); r {
			return k, 1
		} else {
			return k, 0
		}
	}))
}

func Count[K any, V iter.Comparable](i iter.Iterator[K, V], target V) int {
	return CountBy(i, func(v V) bool {
		return v == target
	})
}

func IndexOf[K any, V iter.Comparable](i iter.Iterator[K, V], cmp func(t V) bool) []int {
	res := make([]int, 0, 10)
	idx := 0
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		if cmp(v.V) {
			res = append(res, idx)
		}
		idx++
	}
	return res
}

func NIndexOf[K any, V iter.Comparable](i iter.Iterator[K, V], cmp func(t V) bool, n int) []int {
	res := make([]int, 0, 10)
	idx := 0
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop || len(res) >= n {
			break
		}
		if cmp(v.V) {
			res = append(res, idx)
		}
		idx++
	}
	return res
}
