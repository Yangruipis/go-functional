package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type groupArgs[K iter.Comparable, V any] struct {
	i iter.Iterator[K, V]
	f func(k K, v V) K
}

type groupCases[K iter.Comparable, V any] struct {
	name string
	args groupArgs[K, V]
	want []iter.Entry[K, []V]
}

func testGroupByKey[K iter.Comparable, V any](t *testing.T, cases []groupCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.GroupByKey(tt.args.i)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func testGroupBy[K iter.Comparable, V any](t *testing.T, cases []groupCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.GroupBy(tt.args.i, tt.args.f)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroupByKey(t *testing.T) {
	testGroupByKey(t, []groupCases[int, int]{
		{
			name: "base",
			args: groupArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{K: 0, V: 1},
					{K: 0, V: 2},
					{K: 1, V: 3},
				}),
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2}},
				{K: 1, V: []int{3}},
			},
		},
		{
			name: "empty",
			args: groupArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{}),
			},
			want: []iter.Entry[int, []int]{},
		},
		{
			name: "duplicate k v",
			args: groupArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{K: 0, V: 1},
					{K: 0, V: 2},
					{K: 1, V: 3},
					{K: 1, V: 3},
				}),
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2}},
				{K: 1, V: []int{3, 3}},
			},
		},
	})
}

func TestGroupBy(t *testing.T) {
	testGroupBy(t, []groupCases[int, int]{
		{
			name: "base",
			args: groupArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{K: 0, V: 1},
					{K: 0, V: 2},
					{K: 1, V: 3},
				}),
				f: func(k int, v int) int {
					return k
				},
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2}},
				{K: 1, V: []int{3}},
			},
		},
		{
			name: "by_mod",
			args: groupArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{K: 0, V: 1},
					{K: 1, V: 2},
					{K: 2, V: 3},
				}),
				f: func(k int, v int) int {
					return v % 2
				},
			},
			want: []iter.Entry[int, []int]{
				{K: 1, V: []int{1, 3}},
				{K: 0, V: []int{2}},
			},
		},
	})
}
