package fun

import iter "github.com/Yangruipis/go-functional/pkg/iterator"

func Slice[T1, T2 any](i iter.Iterator[T1, T2]) []T2 {
	res := make([]T2, 0, 100)
	for {
		v, err := i.Next()
		if err == iter.StopIteration {
			break
		}
		res = append(res, v.V)
	}
	return res
}

func ForEach[T1, T2 any](i iter.Iterator[T1, T2], f func(k T1, v T2)) {
	for {
		v, err := i.Next()
		if err == iter.StopIteration {
			break
		}
		f(v.K, v.V)
	}
}
