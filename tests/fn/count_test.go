package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type countArgs[K iter.Comparable, V any] struct {
	i iter.Iterator[K, V]
}

type countCases[K iter.Comparable, V any] struct {
	name string
	args countArgs[K, V]
	want []iter.Entry[K, int]
}

func testCountByKey[K iter.Comparable, V any](t *testing.T, cases []countCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.CountByKey(tt.args.i)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountByKey(t *testing.T) {
	testCountByKey(t, []countCases[int, int]{
		{
			name: "base",
			args: countArgs[int, int]{
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
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 2},
				{K: 1, V: 1},
			},
		},
	})
}

func TestCountByKeyAny(t *testing.T) {
	testCountByKey(t, []countCases[int, interface{}]{
		{
			name: "base",
			args: countArgs[int, interface{}]{
				i: iter.NewEntryIterator([]iter.Entry[int, interface{}]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: "2",
					},
					{
						K: 1,
						V: true,
					},
				}),
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 2},
				{K: 1, V: 1},
			},
		},
	})
}
