package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type aggregateArgs[K iter.Comparable, V, V1 any] struct {
	i iter.Iterator[K, []V]
	f func(vv []V) V1
}

type aggregateCases[K iter.Comparable, V, V1 any] struct {
	name string
	args aggregateArgs[K, V, V1]
	want []iter.Entry[K, V1]
}

func testAggregateBy[K iter.Comparable, V, V1 any](t *testing.T, cases []aggregateCases[K, V, V1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Aggregate(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AggregateByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAggregateByValue(t *testing.T) {
	testAggregateBy(t, []aggregateCases[int, int, int]{
		{
			name: "sum",
			args: aggregateArgs[int, int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, []int]{
					{
						K: 0,
						V: []int{0, 1, 2},
					},
					{
						K: 1,
						V: []int{1, 2, 3},
					},
				}),
				f: func(vv []int) int {
					s := 0
					for _, v := range vv {
						s += v
					}
					return s
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 3},
				{K: 1, V: 6},
			},
		},
	})
}
