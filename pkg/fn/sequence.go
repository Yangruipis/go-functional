package fn

import (
	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

type seq[T1 iter.Comparable, T2 any] struct {
	ScalarIter iter.Iterator[T1, T2]
	SliceIter  iter.Iterator[T1, []T2]

	Paths []string
}

///////////////////////////////////////////////////////////////////////////////
//                               initialization                              //
///////////////////////////////////////////////////////////////////////////////

func Seq[T1 iter.Comparable, T2 any](it iter.Iterator[T1, T2]) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: it,
	}
}

func SliceSeq[T2 any](it []T2) *seq[int, T2] {
	return &seq[int, T2]{
		ScalarIter: iter.NewSliceIterator(it),
	}
}

func MapSeq[T1 iter.Comparable, T2 any](it map[T1]T2) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: iter.NewMapIterator(it),
	}
}

func RangeSeq(start, end, step int) *seq[int, int] {
	return &seq[int, int]{
		ScalarIter: Range(start, end, step),
	}
}

func RepeatSeq[T any](t T, num int) *seq[int, T] {
	return &seq[int, T]{
		ScalarIter: Repeat(t, num),
	}
}

///////////////////////////////////////////////////////////////////////////////
//                               transofrmation                              //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[T1, T2]) Map(f func(inK T1, inV T2) (T1, T2)) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Map(s.getScalarIter(), f),
		Paths:      append(s.Paths, "Map"),
	}
}

func (s *seq[T1, T2]) Filter(f func(inK T1, inV T2) bool) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Filter(s.getScalarIter(), f),
		Paths:      append(s.Paths, "Filter"),
	}
}

func (s *seq[T1, T2]) Flatten() *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Flatten(s.getSliceIter()),
		Paths:      append(s.Paths, "Flatten"),
	}
}

func (s *seq[T1, T2]) GroupByKey() *seq[T1, T2] {
	return &seq[T1, T2]{
		SliceIter: GroupByKey(s.getScalarIter()),
		Paths:     append(s.Paths, "GroupByKey"),
	}
}

func (s *seq[T1, T2]) GroupBy(f func(k T1, v T2) T1) *seq[T1, T2] {
	return &seq[T1, T2]{
		SliceIter: GroupBy(s.getScalarIter(), f),
		Paths:     append(s.Paths, "GroupBy"),
	}
}

func (s *seq[T1, T2]) FlatMap(f func(k T1, v []T2) (T1, []T2)) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: FlatMap(s.getSliceIter(), f),
		Paths:      append(s.Paths, "FlatMap"),
	}
}

func (s *seq[T1, T2]) ReduceByKey(f func(a, b T2) T2) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: ReduceByKey(s.getScalarIter(), f),
		Paths:      append(s.Paths, "ReduceByKey"),
	}
}

func (s *seq[T1, T2]) CountByKey(f func(a, b T2) T2) *seq[T1, int] {
	return &seq[T1, int]{
		ScalarIter: CountByKey(s.getScalarIter()),
		Paths:      append(s.Paths, "CountByKey"),
	}
}

func (s *seq[T1, T2]) Reverse(f func(a, b T2) T2) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Reverse(s.getScalarIter()),
		Paths:      append(s.Paths, "Reverse"),
	}
}

func (s *seq[T1, T2]) Chunk(size int) *seq[int, T2] {
	return &seq[int, T2]{
		SliceIter: Chunk(s.getScalarIter(), size),
		Paths:     append(s.Paths, "Chunk"),
	}
}

func (s *seq[T1, T2]) Sort(lessFn func(v1, v2 T2) bool) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Sort(s.getScalarIter(), lessFn),
		Paths:      append(s.Paths, "Sort"),
	}
}

func (s *seq[T1, T2]) Aggregate(f func(vv []T2) T2) *seq[T1, T2] {

	return &seq[T1, T2]{
		ScalarIter: Aggregate(s.getSliceIter(), f),
		Paths:      append(s.Paths, "Aggregate"),
	}
}

func (s *seq[T1, T2]) Shuffle() *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Shuffle(s.getScalarIter()),
		Paths:      append(s.Paths, "Shuffle"),
	}
}

func (s *seq[T1, T2]) Choices(size float32) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Choices(s.getScalarIter(), size),
		Paths:      append(s.Paths, "Choices"),
	}
}

func (s *seq[T1, T2]) Sample(size float32) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Sample(s.getScalarIter(), size),
		Paths:      append(s.Paths, "Sample"),
	}
}

func (s *seq[T1, T2]) Head(n int) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Head(s.getScalarIter(), n),
		Paths:      append(s.Paths, "Head"),
	}
}

func (s *seq[T1, T2]) Tail(n int) *seq[T1, T2] {
	return &seq[T1, T2]{
		ScalarIter: Tail(s.getScalarIter(), n),
		Paths:      append(s.Paths, "Tail"),
	}
}

///////////////////////////////////////////////////////////////////////////////
//                                   action                                  //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[T1, T2]) ToSlice() []T2 {
	return ToSlice(s.getScalarIter())
}

func (s *seq[T1, T2]) ToMap() map[T1]T2 {
	return ToMap(s.getScalarIter())
}

func (s *seq[T1, T2]) Keys() []T1 {
	return Keys(s.getScalarIter())
}

func (s *seq[T1, T2]) Values() []T2 {
	return Values(s.getScalarIter())
}

func (s *seq[T1, T2]) Entries() []iter.Entry[T1, T2] {
	return Entries(s.getScalarIter())
}

func (s *seq[T1, T2]) ForEach(f func(k T1, v T2)) {
	ForEach(s.getScalarIter(), f)
}

func (s *seq[T1, T2]) Reduce(f func(a, b T2) T2) T2 {
	return Reduce(s.getScalarIter(), f)
}

func (s *seq[T1, T2]) Size() int {
	return Size(s.getScalarIter())
}

func (s *seq[T1, T2]) Any(f func(T1, T2) (T1, bool)) bool {
	return Any(Map(s.getScalarIter(), f))
}

func (s *seq[T1, T2]) All(f func(T1, T2) (T1, bool)) bool {
	return All(Map(s.getScalarIter(), f))
}

func (s *seq[T1, T2]) ExistsBy(f func(T2) bool) bool {
	return ExistsBy(s.getScalarIter(), f)
}

func (s *seq[T1, T2]) CountBy(f func(T2) bool) int {
	return CountBy(s.getScalarIter(), f)
}

///////////////////////////////////////////////////////////////////////////////
//                                   utils                                   //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[T1, T2]) getScalarIter() iter.Iterator[T1, T2] {
	if s.ScalarIter == nil {
		panic("scalar iterator is nil, make sure your value's type is T2")
	}
	return s.ScalarIter
}

func (s *seq[T1, T2]) getSliceIter() iter.Iterator[T1, []T2] {
	if s.SliceIter == nil {
		panic("slice iterator is nil, make sure your value's type is []T2")
	}
	return s.SliceIter
}
