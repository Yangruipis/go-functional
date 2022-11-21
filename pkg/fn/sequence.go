package fn

import (
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type seq[K iter.Comparable, V any] struct {
	ScalarIter iter.Iterator[K, V]
	SliceIter  iter.Iterator[K, []V]

	Paths []string
}

///////////////////////////////////////////////////////////////////////////////
//                               initialization                              //
///////////////////////////////////////////////////////////////////////////////

func Seq[K iter.Comparable, V any](it iter.Iterator[K, V]) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: it,
	}
}

func SliceSeq[V any](it []V) *seq[int, V] {
	return &seq[int, V]{
		ScalarIter: iter.NewSliceIterator(it),
	}
}

func MapSeq[K iter.Comparable, V any](it map[K]V) *seq[K, V] {
	return &seq[K, V]{
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
//                               transformation                              //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[K, V]) Map(f func(inK K, inV V) (K, V)) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Map(s.getScalarIter(), f),
		Paths:      append(s.Paths, "Map"),
	}
}

func (s *seq[K, V]) Filter(f func(inK K, inV V) bool) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Filter(s.getScalarIter(), f),
		Paths:      append(s.Paths, "Filter"),
	}
}

func (s *seq[K, V]) Flatten() *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Flatten(s.getSliceIter()),
		Paths:      append(s.Paths, "Flatten"),
	}
}

func (s *seq[K, V]) GroupByKey() *seq[K, V] {
	return &seq[K, V]{
		SliceIter: GroupByKey(s.getScalarIter()),
		Paths:     append(s.Paths, "GroupByKey"),
	}
}

func (s *seq[K, V]) GroupBy(f func(k K, v V) K) *seq[K, V] {
	return &seq[K, V]{
		SliceIter: GroupBy(s.getScalarIter(), f),
		Paths:     append(s.Paths, "GroupBy"),
	}
}

func (s *seq[K, V]) FlatMap(f func(k K, v []V) (K, []V)) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: FlatMap(s.getSliceIter(), f),
		Paths:      append(s.Paths, "FlatMap"),
	}
}

func (s *seq[K, V]) ReduceByKey(f func(a, b V) V) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: ReduceByKey(s.getScalarIter(), f),
		Paths:      append(s.Paths, "ReduceByKey"),
	}
}

func (s *seq[K, V]) CountByKey(f func(a, b V) V) *seq[K, int] {
	return &seq[K, int]{
		ScalarIter: CountByKey(s.getScalarIter()),
		Paths:      append(s.Paths, "CountByKey"),
	}
}

func (s *seq[K, V]) Reverse(f func(a, b V) V) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Reverse(s.getScalarIter()),
		Paths:      append(s.Paths, "Reverse"),
	}
}

func (s *seq[K, V]) Chunk(size int) *seq[int, V] {
	return &seq[int, V]{
		SliceIter: Chunk(s.getScalarIter(), size),
		Paths:     append(s.Paths, "Chunk"),
	}
}

func (s *seq[K, V]) Sort(lessFn func(v1, v2 V) bool) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Sort(s.getScalarIter(), lessFn),
		Paths:      append(s.Paths, "Sort"),
	}
}

func (s *seq[K, V]) Aggregate(f func(vv []V) V) *seq[K, V] {

	return &seq[K, V]{
		ScalarIter: Aggregate(s.getSliceIter(), f),
		Paths:      append(s.Paths, "Aggregate"),
	}
}

func (s *seq[K, V]) Shuffle() *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Shuffle(s.getScalarIter()),
		Paths:      append(s.Paths, "Shuffle"),
	}
}

func (s *seq[K, V]) Choices(size float32) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Choices(s.getScalarIter(), size),
		Paths:      append(s.Paths, "Choices"),
	}
}

func (s *seq[K, V]) Sample(size float32) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Sample(s.getScalarIter(), size),
		Paths:      append(s.Paths, "Sample"),
	}
}

func (s *seq[K, V]) Head(n int) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Head(s.getScalarIter(), n),
		Paths:      append(s.Paths, "Head"),
	}
}

func (s *seq[K, V]) Tail(n int) *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Tail(s.getScalarIter(), n),
		Paths:      append(s.Paths, "Tail"),
	}
}

func (s *seq[K, V]) Cache() *seq[K, V] {
	return &seq[K, V]{
		ScalarIter: Cache(s.ScalarIter),
		Paths:      append(s.Paths, "Cache"),
	}
}

///////////////////////////////////////////////////////////////////////////////
//                                   action                                  //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[K, V]) ToSlice() []V {
	return ToSlice(s.getScalarIter())
}

func (s *seq[K, V]) ToMap() map[K]V {
	return ToMap(s.getScalarIter())
}

func (s *seq[K, V]) Keys() []K {
	return Keys(s.getScalarIter())
}

func (s *seq[K, V]) Values() []V {
	return Values(s.getScalarIter())
}

func (s *seq[K, V]) Entries() []iter.Entry[K, V] {
	return Entries(s.getScalarIter())
}

func (s *seq[K, V]) ForEach(f func(k K, v V)) {
	ForEach(s.getScalarIter(), f)
}

func (s *seq[K, V]) Reduce(f func(a, b V) V) V {
	return Reduce(s.getScalarIter(), f)
}

func (s *seq[K, V]) Size() int {
	return Size(s.getScalarIter())
}

func (s *seq[K, V]) Any(f func(K, V) (K, bool)) bool {
	return Any(Map(s.getScalarIter(), f))
}

func (s *seq[K, V]) All(f func(K, V) (K, bool)) bool {
	return All(Map(s.getScalarIter(), f))
}

func (s *seq[K, V]) ExistsBy(f func(V) bool) bool {
	return ExistsBy(s.getScalarIter(), f)
}

func (s *seq[K, V]) CountBy(f func(V) bool) int {
	return CountBy(s.getScalarIter(), f)
}

///////////////////////////////////////////////////////////////////////////////
//                                   utils                                   //
///////////////////////////////////////////////////////////////////////////////

func (s *seq[K, V]) getScalarIter() iter.Iterator[K, V] {
	if s.ScalarIter == nil {
		panic("scalar iterator is nil, make sure your value's type is V")
	}
	return s.ScalarIter
}

func (s *seq[K, V]) getSliceIter() iter.Iterator[K, []V] {
	if s.SliceIter == nil {
		panic("slice iterator is nil, make sure your value's type is []V")
	}
	return s.SliceIter
}
