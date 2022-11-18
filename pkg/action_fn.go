package fun

import (
	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

func Slice[T1, T2 any](i iter.Iterator[T1, T2]) []T2 {
	res := make([]T2, 0, 100)
	for {
		v, flag := i.Next()
		if flag == iter.FlagStop {
			break
		}
		res = append(res, v.V)
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
