package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type reduceArgs[K iter.Comparable, V any] struct {
	i iter.Iterator[K, V]
	f func(a, b V) V
}

type reduceCases[K iter.Comparable, V any] struct {
	name string
	args reduceArgs[K, V]
	want []iter.Entry[K, V]
}

func testReduceByKey[K iter.Comparable, V any](t *testing.T, cases []reduceCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.ReduceByKey(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReduceByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReduceByKey(t *testing.T) {
	testReduceByKey(t, []reduceCases[int, int]{
		{
			name: "base",
			args: reduceArgs[int, int]{
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
				f: func(a, b int) int {
					return a + b
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 3},
				{K: 1, V: 1},
			},
		},
	})
}
