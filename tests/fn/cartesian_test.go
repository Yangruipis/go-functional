package fn_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type cartesianArgs[K, V, V1 any] struct {
	i iter.Iterator[K, V]
	j iter.Iterator[K, V]
	f func(v1, v2 V) V1
}

type cartesianCases[K, V, V1 any] struct {
	name string
	args cartesianArgs[K, V, V1]
	want []iter.Entry[K, V1]
}

func testCartesianBy[K, V, V1 any](t *testing.T, cases []cartesianCases[K, V, V1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Cartesian(tt.args.i, tt.args.j, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartesianByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCartesianByValue(t *testing.T) {
	testCartesianBy(t, []cartesianCases[int, int, int]{
		{
			name: "byval",
			args: cartesianArgs[int, int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 3,
					},
				}),
				j: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 2,
					},
				}),
				f: func(v1, v2 int) int {
					return v1 * v2
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 1, V: 2},
				{K: 2, V: 3},
				{K: 3, V: 6},
			},
		},
	})
}

func TestCartesianByValue2(t *testing.T) {
	testCartesianBy(t, []cartesianCases[int, int, string]{
		{
			name: "byval",
			args: cartesianArgs[int, int, string]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 3,
					},
				}),
				j: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 2,
					},
				}),
				f: func(v1, v2 int) string {
					return fmt.Sprintf("%d%d", v1, v2)
				},
			},
			want: []iter.Entry[int, string]{
				{K: 0, V: "11"},
				{K: 1, V: "12"},
				{K: 2, V: "31"},
				{K: 3, V: "32"},
			},
		},
	})
}
