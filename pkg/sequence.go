package fun

import (
	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

type seq[T1 any, T2 any] struct {
	Iter iter.Iterator[T1, T2]
}

func Seq[T1, T2 any](it iter.Iterator[T1, T2]) *seq[T1, T2] {
	return &seq[T1, T2]{
		Iter: it,
	}
}

func SliceSeq[T2 any](it []T2) *seq[int, T2] {
	return &seq[int, T2]{
		Iter: iter.NewSliceIterator(it),
	}
}

func MapSeq[T1 iter.Comparable, T2 any](it map[T1]T2) *seq[T1, T2] {
	return &seq[T1, T2]{
		Iter: iter.NewMapIterator(it),
	}
}

func RangeSeq(start, end, step int) *seq[int, int] {
	return &seq[int, int]{
		Iter: Range(start, end, step),
	}
}

func RepeatSeq[T any](t T, num int) *seq[int, T] {
	return &seq[int, T]{
		Iter: Repeat(t, num),
	}
}

func (s *seq[T1, T2]) Map(f func(inK T1, inV T2) (T1, T2)) *seq[T1, T2] {
	return &seq[T1, T2]{
		Iter: Map(s.Iter, f),
	}
}

func (s *seq[T1, T2]) Filter(f func(inK T1, inV T2) bool) *seq[T1, T2] {
	return &seq[T1, T2]{
		Iter: Filter(s.Iter, f),
	}
}

func (s *seq[T1, T2]) ToSlice() []T2 {
	return ToSlice(s.Iter)
}

func (s *seq[T1, T2]) ForEach(f func(k T1, v T2)) {
	ForEach(s.Iter, f)
}

func (s *seq[T1, T2]) Reduce(f func(a, b T2) T2) T2 {
	return Reduce(s.Iter, f)
}
