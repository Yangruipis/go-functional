package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type sortArgs[K, V any] struct {
	i iter.Iterator[K, V]
	f func(v1, v2 V) bool
}

type sortCases[K, V any] struct {
	name string
	args sortArgs[K, V]
	want []iter.Entry[K, V]
}

func testSortBy[K, V any](t *testing.T, cases []sortCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Sort(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	testSortBy(t, []sortCases[int, int]{
		{
			name: "byval",
			args: sortArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 3,
					},
					{
						K: 2,
						V: 2,
					},
				}),
				f: func(v1, v2 int) bool {
					return v1 < v2
				},
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 2, V: 2},
				{K: 1, V: 3},
			},
		},
	})
}

func TestSortReversed(t *testing.T) {
	testSortBy(t, []sortCases[int, int]{
		{
			name: "reversed",
			args: sortArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 3,
					},
					{
						K: 2,
						V: 2,
					},
				}),
				f: func(v1, v2 int) bool {
					return v1 > v2
				},
			},
			want: []iter.Entry[int, int]{
				{K: 1, V: 3},
				{K: 2, V: 2},
				{K: 0, V: 1},
			},
		},
	})
}

func TestSortStable(t *testing.T) {
	testSortBy(t, []sortCases[int, int]{
		{
			name: "stable",
			args: sortArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 3,
					},
					{
						K: 2,
						V: 3,
					},
					{
						K: 3,
						V: 3,
					},
					{
						K: 4,
						V: 2,
					},
				}),
				f: func(v1, v2 int) bool {
					return v1 > v2
				},
			},
			want: []iter.Entry[int, int]{
				{K: 1, V: 3},
				{K: 2, V: 3},
				{K: 3, V: 3},
				{K: 4, V: 2},
				{K: 0, V: 1},
			},
		},
	})
}
