package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type intersectArgs[K, V any, K1 iter.Comparable] struct {
	i iter.Iterator[K, V]
	j iter.Iterator[K, V]
	f func(k K, v V) K1
}

type intersectCases[K, V any, K1 iter.Comparable] struct {
	name string
	args intersectArgs[K, V, K1]
	want []iter.Entry[K, V]
}

func testIntersectBy[K, V any, K1 iter.Comparable](t *testing.T, cases []intersectCases[K, V, K1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.IntersectBy(tt.args.i, tt.args.j, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectByValue(t *testing.T) {
	testIntersectBy(t, []intersectCases[int, int, int]{
		{
			name: "byval",
			args: intersectArgs[int, int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 2,
					},
					{
						K: 1,
						V: 1,
					},
				}),
				j: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 3,
					},
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 4,
					},
				}),
				f: func(k, v int) int {
					return v
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 1, V: 1},
			},
		},
	})
}

func TestIntersectByByKey(t *testing.T) {
	testIntersectBy(t, []intersectCases[int, int, int]{
		{
			name: "bykey",
			args: intersectArgs[int, int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 2,
					},
					{
						K: 1,
						V: 1,
					},
				}),
				j: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 3,
					},
					{
						K: 0,
						V: 0,
					},
					{
						K: 1,
						V: 4,
					},
				}),
				f: func(k, v int) int {
					return k
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 0, V: 2},
				{K: 1, V: 1},
			},
		},
	})
}
