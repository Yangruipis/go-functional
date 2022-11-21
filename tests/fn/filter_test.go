package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type filterArgs[K, V any] struct {
	i iter.Iterator[K, V]
	f func(k K, v V) bool
}

type filterCases[K, V any] struct {
	name string
	args filterArgs[K, V]
	want []iter.Entry[K, V]
}

func testFilter[K, V any](t *testing.T, cases []filterCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Filter(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterIntSlice(t *testing.T) {
	testFilter(t, []filterCases[int, int]{
		{
			name: "base",
			args: filterArgs[int, int]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) bool {
					return v >= 2
				},
			},
			want: []iter.Entry[int, int]{
				{K: 1, V: 2},
				{K: 2, V: 3},
			},
		},
		{
			name: "empty",
			args: filterArgs[int, int]{
				i: iter.NewSliceIterator([]int{1, 2, 3}),
				f: func(k int, v int) bool {
					return v >= 4
				},
			},
			want: []iter.Entry[int, int]{},
		},
	})
}
