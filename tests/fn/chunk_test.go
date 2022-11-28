package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type chunkArgs[K, V any] struct {
	i iter.Iterator[K, V]
	n int
}

type chunkCases[K, V any] struct {
	name string
	args chunkArgs[K, V]
	want []iter.Entry[K, []V]
}

func testChunkByKey[K iter.Comparable, V any](t *testing.T, cases []chunkCases[K, V]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Chunk(tt.args.i, tt.args.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChunkByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunkByKey(t *testing.T) {
	testChunkByKey(t, []chunkCases[int, int]{
		{
			name: "base",
			args: chunkArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 2,
					},
					{
						K: 2,
						V: 3,
					},
				}),
				n: 1,
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1}},
				{K: 1, V: []int{2}},
				{K: 2, V: []int{3}},
			},
		},
		{
			name: "chunk2",
			args: chunkArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 2,
					},
					{
						K: 2,
						V: 3,
					},
				}),
				n: 2,
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2}},
				{K: 1, V: []int{3}},
			},
		},
		{
			name: "chunk3",
			args: chunkArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 2,
					},
					{
						K: 2,
						V: 3,
					},
				}),
				n: 3,
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2, 3}},
			},
		},
		{
			name: "chunk4",
			args: chunkArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 1,
						V: 2,
					},
					{
						K: 2,
						V: 3,
					},
				}),
				n: 4,
			},
			want: []iter.Entry[int, []int]{
				{K: 0, V: []int{1, 2, 3}},
			},
		},
		{
			name: "chunk_empty",
			args: chunkArgs[int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{}),
				n: 1,
			},
			want: []iter.Entry[int, []int]{},
		},
	})
}
