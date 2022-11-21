package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	iter "github.com/Yangruipis/go-functional/pkg/iterator"
)

type mapArgs[K, V, K1, V1 any] struct {
	i iter.Iterator[K, V]
	f func(k K, v V) (K1, V1)
}

type mapCases[K, V, K1, V1 any] struct {
	name string
	args mapArgs[K, V, K1, V1]
	want []iter.Entry[K1, V1]
}

func testMap[K, V, K1, V1 any](t *testing.T, cases []mapCases[K, V, K1, V1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Map(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapIntInt(t *testing.T) {
	testMap(t, []mapCases[int, int, int, int]{
		{
			name: "base",
			args: mapArgs[int, int, int, int]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) (int, int) {
					return k, v + 1
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 2},
				{K: 1, V: 3},
				{K: 2, V: 4},
			},
		},
	})
}

func TestMapIntStruct(t *testing.T) {
	type out struct {
		num int
	}
	testMap(t, []mapCases[int, int, int, out]{
		{
			name: "base",
			args: mapArgs[int, int, int, out]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) (int, out) {
					return k, out{
						num: v,
					}
				},
			},
			want: []iter.Entry[int, out]{
				{K: 0, V: out{num: 1}},
				{K: 1, V: out{num: 2}},
				{K: 2, V: out{num: 3}},
			},
		},
	})
}

func TestMapIntStructPointer(t *testing.T) {
	type out struct {
		num int
	}
	testMap(t, []mapCases[int, int, int, *out]{
		{
			name: "base",
			args: mapArgs[int, int, int, *out]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) (int, *out) {
					return k, &out{
						num: v,
					}
				},
			},
			want: []iter.Entry[int, *out]{
				{K: 0, V: &out{num: 1}},
				{K: 1, V: &out{num: 2}},
				{K: 2, V: &out{num: 3}},
			},
		},
		{
			name: "with_nil",
			args: mapArgs[int, int, int, *out]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) (int, *out) {
					if k == 1 {
						return k, &out{
							num: v,
						}
					}
					return k, nil
				},
			},
			want: []iter.Entry[int, *out]{
				{K: 0, V: nil},
				{K: 1, V: &out{num: 2}},
				{K: 2, V: nil},
			},
		},
	})
}
