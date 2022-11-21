package fn

import (
	"reflect"
	"testing"

	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

type K int
type V int

type K1 int
type V1 int

func TestFilter(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		f func(k K, v V) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatten(t *testing.T) {
	type args struct {
		i iter.Iterator[K, []V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flatten(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByKey(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, []V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByKey(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		f func(k K, v V) K
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, []V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupBy(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlatMap(t *testing.T) {
	type args struct {
		i iter.Iterator[K, []V]
		f func(k K, v []V) (K, []V)
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatMap(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIter(t *testing.T) {
	type args struct {
		arr []V
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIter(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIter(t *testing.T) {
	type args struct {
		m map[K]V
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapIter(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapIter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRange(t *testing.T) {
	type args struct {
		start int
		end   int
		step  int
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Range(tt.args.start, tt.args.end, tt.args.step); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepeat(t *testing.T) {
	type args struct {
		t   V
		num int
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.t, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceByKey(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		f func(a, b V) V
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReduceByKey(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReduceByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountByKey(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountByKey(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[V, K]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByVal(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[V, []K]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GroupByVal(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GroupByVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountByVal(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[V, int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountByVal(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountByVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionBy(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
		f  func(k K, v V) K1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnionBy(tt.args.i1, tt.args.i2, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectBy(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
		f  func(k K, v V) K1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectBy(tt.args.i1, tt.args.i2, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstractBy(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
		f  func(k K, v V) K1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SubstractBy(tt.args.i1, tt.args.i2, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubstractBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctBy(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		f func(k K, v V) K1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DistinctBy(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnion(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Union(tt.args.i1, tt.args.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.i1, tt.args.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.i1, tt.args.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinct(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distinct(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Distinct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartesian(t *testing.T) {
	type args struct {
		i1 iter.Iterator[K, V]
		i2 iter.Iterator[K, V]
		f  func(v1, v2 V) V1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, V1]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cartesian(tt.args.i1, tt.args.i2, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cartesian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args struct {
		i    iter.Iterator[K, V]
		size int
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, []V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		i      iter.Iterator[K, V]
		lessFn func(v1, v2 V) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.i, tt.args.lessFn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregate(t *testing.T) {
	type args struct {
		i iter.Iterator[K, []V]
		f func(vv []V) V1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V1]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Aggregate(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip(t *testing.T) {
	type args struct {
		i1 iter.Iterator[int, V]
		i2 iter.Iterator[int, V1]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[int, iter.Entry[V1, V1]]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zip(tt.args.i1, tt.args.i2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShuffle(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Shuffle(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shuffle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSample(t *testing.T) {
	type args struct {
		i    iter.Iterator[K, V]
		size float32
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sample(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChoices(t *testing.T) {
	type args struct {
		i    iter.Iterator[K, V]
		size float32
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Choices(tt.args.i, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Choices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHead(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		n int
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Head(tt.args.i, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Head() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTail(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
		n int
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Tail(tt.args.i, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCache(t *testing.T) {
	type args struct {
		i iter.Iterator[K, V]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[K, V]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Cache(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cache() = %v, want %v", got, tt.want)
			}
		})
	}
}
