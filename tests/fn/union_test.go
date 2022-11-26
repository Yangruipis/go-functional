package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type unionArgs[K, V any, K1 iter.Comparable] struct {
	i iter.Iterator[K, V]
	j iter.Iterator[K, V]
	f func(k K, v V) K1
}

type unionCases[K, V any, K1 iter.Comparable] struct {
	name string
	args unionArgs[K, V, K1]
	want []iter.Entry[K, V]
}

func testUnionBy[K, V any, K1 iter.Comparable](t *testing.T, cases []unionCases[K, V, K1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.UnionBy(tt.args.i, tt.args.j, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionByValue(t *testing.T) {
	testUnionBy(t, []unionCases[int, int, int]{
		{
			name: "byval",
			args: unionArgs[int, int, int]{
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
					return v
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 0, V: 2},
				{K: 0, V: 3},
				{K: 0, V: 0},
				{K: 1, V: 4},
			},
		},
	})
}

func TestUnionByByKey(t *testing.T) {
	testUnionBy(t, []unionCases[int, int, int]{
		{
			name: "bykey",
			args: unionArgs[int, int, int]{
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
				{K: 1, V: 1},
			},
		},
	})
}
