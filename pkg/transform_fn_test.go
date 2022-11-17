package fun

import (
	"reflect"
	"testing"

	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

type (
	T1 int
	T2 int
	O1 int
	O2 int
)

func TestMap(t *testing.T) {
	type args struct {
		i iter.Iterator[T1, T2]
		f func(k T1, v T2) (O1, O2)
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[O1, O2]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.i, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		i iter.Iterator[T1, T2]
		f func(k T1, v T2) bool
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[T1, T2]
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
		i iter.Iterator[T1, []T2]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[T1, T2]
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
		i iter.Iterator[T1, T2]
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[T1, []T2]
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
		i iter.Iterator[T1, T2]
		f func(k T1) T1
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[T1, []T2]
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
		i iter.Iterator[T1, []T2]
		f func(k T1, v []T2) (T1, []T2)
	}
	tests := []struct {
		name string
		args args
		want iter.Iterator[T1, T2]
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
