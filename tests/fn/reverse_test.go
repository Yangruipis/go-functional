package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type invertArgs[K, V any] struct {
	i iter.Iterator[K, V]
}

type invertCases[K, V any] struct {
	name string
	args invertArgs[K, V]
	want []iter.Entry[V, K]
}

func testInvert[K, V any](t *testing.T, cases []invertCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Invert(tt.args.i)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvert(t *testing.T) {
	testInvert(t, []invertCases[int, bool]{
		{
			name: "base",
			args: invertArgs[int, bool]{
				i: iter.NewEntryIterator([]iter.Entry[int, bool]{
					{
						K: 0,
						V: true,
					},
					{
						K: 1,
						V: false,
					},
					{
						K: 1,
						V: true,
					},
				}),
			},
			want: []iter.Entry[bool, int]{
				{K: true, V: 0},
				{K: false, V: 1},
				{K: true, V: 1},
			},
		},
	})
}

type reverseArgs[K, V any] struct {
	i iter.Iterator[K, V]
}

type reverseCases[K, V any] struct {
	name string
	args reverseArgs[K, V]
	want []iter.Entry[K, V]
}

func testReverse[K, V any](t *testing.T, cases []reverseCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Reverse(tt.args.i)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testReverse(t, []reverseCases[int, bool]{
		{
			name: "base",
			args: reverseArgs[int, bool]{
				i: iter.NewEntryIterator([]iter.Entry[int, bool]{
					{
						K: 0,
						V: true,
					},
					{
						K: 1,
						V: false,
					},
					{
						K: 1,
						V: true,
					},
				}),
			},
			want: []iter.Entry[int, bool]{
				{K: 1, V: true},
				{K: 1, V: false},
				{K: 0, V: true},
			},
		},
	})
}
