package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type flattenArgs[K, V any] struct {
	i iter.Iterator[K, []V]
}

type flattenCases[K, V any] struct {
	name string
	args flattenArgs[K, V]
	want []iter.Entry[K, V]
}

func testFlatten[K, V any](t *testing.T, cases []flattenCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Flatten(tt.args.i)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flatten() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlattenIntSlice(t *testing.T) {
	testFlatten(t, []flattenCases[int, int]{
		{
			name: "base",
			args: flattenArgs[int, int]{
				i: iter.NewSliceIterator([][]int{{1, 2, 3}, {4, 5}}),
			},
			want: []iter.Entry[int, int]{
				{K: 0, V: 1},
				{K: 0, V: 2},
				{K: 0, V: 3},
				{K: 1, V: 4},
				{K: 1, V: 5},
			},
		},
		{
			name: "empty",
			args: flattenArgs[int, int]{
				i: iter.NewSliceIterator([][]int{{}, {4, 5}}),
			},
			want: []iter.Entry[int, int]{
				{K: 1, V: 4},
				{K: 1, V: 5},
			},
		},
		{
			name: "empty",
			args: flattenArgs[int, int]{
				i: iter.NewSliceIterator([][]int{}),
			},
			want: []iter.Entry[int, int]{},
		},
	})
}

func TestFlattenStrMap(t *testing.T) {
	testFlatten(t, []flattenCases[string, int]{
		{
			name: "base",
			args: flattenArgs[string, int]{
				i: iter.NewMapIterator(map[string][]int{
					"a": {1, 2},
					"b": {3},
				}),
			},
			want: []iter.Entry[string, int]{
				{K: "a", V: 1},
				{K: "a", V: 2},
				{K: "b", V: 3},
			},
		},
	})
}

func TestFlattenNested(t *testing.T) {
	testFlatten(t, []flattenCases[string, []int]{
		{
			name: "base",
			args: flattenArgs[string, []int]{
				i: iter.NewMapIteratorWithKeys([]string{"a", "b"}, map[string][][]int{
					"a": {{1}, {2}},
					"b": {{3}},
				}),
			},
			want: []iter.Entry[string, []int]{
				{K: "a", V: []int{1}},
				{K: "a", V: []int{2}},
				{K: "b", V: []int{3}},
			},
		},
	})
}
