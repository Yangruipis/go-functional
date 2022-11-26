package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type distinctArgs[K, V any, K1 iter.Comparable] struct {
	i iter.Iterator[K, V]
	f func(k K, v V) K1
}

type distinctCases[K, V any, K1 iter.Comparable] struct {
	name string
	args distinctArgs[K, V, K1]
	want []iter.Entry[K, V]
}

func testDistinctBy[K, V any, K1 iter.Comparable](t *testing.T, cases []distinctCases[K, V, K1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.DistinctBy(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DistinctByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistinctByValue(t *testing.T) {
	testDistinctBy(t, []distinctCases[int, int, int]{
		{
			name: "byval",
			args: distinctArgs[int, int, int]{
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
				f: func(k, v int) int {
					return v
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 0, V: 2},
			},
		},
	})
}

func TestDistinctByByKey(t *testing.T) {
	testDistinctBy(t, []distinctCases[int, int, int]{
		{
			name: "bykey",
			args: distinctArgs[int, int, int]{
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
