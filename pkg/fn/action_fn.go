package fn

import (
	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

func ToSlice[T1, T2 any](i iter.Iterator[T1, T2]) []T2 {
	res := make([]T2, 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v.V)
	}
	return res
}

func ToMap[T1 iter.Comparable, T2 any](i iter.Iterator[T1, T2]) map[T1]T2 {
	r := make(map[T1]T2)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		r[v.K] = v.V
	}
	return r
}

func ToSet[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2]) map[T2]struct{} {
	r := make(map[T2]struct{})
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		r[v.V] = struct{}{}
	}
	return r
}

func Values[T1, T2 any](i iter.Iterator[T1, T2]) []T2 {
	return ToSlice(i)
}

func Keys[T1, T2 any](i iter.Iterator[T1, T2]) []T1 {
	res := make([]T1, 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v.K)
	}
	return res
}

func Entries[T1, T2 any](i iter.Iterator[T1, T2]) []iter.Entry[T1, T2] {
	res := make([]iter.Entry[T1, T2], 0, 16)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v)
	}
	return res
}

func ForEach[T1, T2 any](i iter.Iterator[T1, T2], f func(k T1, v T2)) {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		f(v.K, v.V)
	}
}

func Reduce[T1, T2 any](i iter.Iterator[T1, T2], f func(a, b T2) T2) T2 {
	var (
		res T2
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

func Size[T1, T2 any](i iter.Iterator[T1, T2]) int {
	return len(ToSlice(i))
}

func Any[T1 any, T2 bool](i iter.Iterator[T1, T2]) bool {
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

func All[T1 any, T2 bool](i iter.Iterator[T1, T2]) bool {
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

func ExistsBy[T1, T2 any](i iter.Iterator[T1, T2], eq func(vsrc T2) bool) bool {
	return Any(Map(i, func(k T1, v T2) (T1, bool) {
		return k, eq(v)
	}))
}

func Exists[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2], target T2) bool {
	return ExistsBy(i, func(v T2) bool {
		return v == target
	})
}

func Contains[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2], target T2) bool {
	return Exists(i, target)
}

func Sum[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2]) (res T2) {
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res += v.V
	}
	return
}

func Avg[T1 any, T2 iter.Numeric | iter.Int | iter.Uint](i iter.Iterator[T1, T2]) (res T2) {
	cnt := 0
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res += v.V
		cnt++
	}
	return res / T2(cnt)
}

func CountBy[T1, T2 any](i iter.Iterator[T1, T2], cmp func(v T2) bool) int {
	return Sum(Map(i, func(k T1, v T2) (T1, int) {
		if r := cmp(v); r {
			return k, 1
		} else {
			return k, 0
		}
	}))
}

func Count[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2], target T2) int {
	return CountBy(i, func(v T2) bool {
		return v == target
	})
}

func IndexOf[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2], cmp func(t T2) bool) []int {
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

func NIndexOf[T1 any, T2 iter.Comparable](i iter.Iterator[T1, T2], cmp func(t T2) bool, n int) []int {
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
