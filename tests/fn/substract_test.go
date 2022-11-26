package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type substractArgs[K, V any, K1 iter.Comparable] struct {
	i iter.Iterator[K, V]
	j iter.Iterator[K, V]
	f func(k K, v V) K1
}

type substractCases[K, V any, K1 iter.Comparable] struct {
	name string
	args substractArgs[K, V, K1]
	want []iter.Entry[K, V]
}

func testSubstractBy[K, V any, K1 iter.Comparable](t *testing.T, cases []substractCases[K, V, K1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.SubstractBy(tt.args.i, tt.args.j, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubstractByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstractByValue(t *testing.T) {
	testSubstractBy(t, []substractCases[int, int, int]{
		{
			name: "byval",
			args: substractArgs[int, int, int]{
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
				{K: 0, V: 2},
			},
		},
	})
}

func TestSubstractByByKey(t *testing.T) {
	testSubstractBy(t, []substractCases[int, int, int]{
		{
			name: "bykey",
			args: substractArgs[int, int, int]{
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
			want: []iter.Entry[int, int]{},
		},
	})
}
