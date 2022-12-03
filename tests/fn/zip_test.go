package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

type zipArgs[K int, V, V1 any] struct {
	i iter.Iterator[K, V]
	j iter.Iterator[K, V1]
}

type zipCases[K int, V, V1 any] struct {
	name string
	args zipArgs[K, V, V1]
	want []iter.Entry[K, iter.Entry[V, V1]]
}

func testZipBy[V, V1 any](t *testing.T, cases []zipCases[int, V, V1]) {
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			if got := fn.Entries(fn.Zip(tt.args.i, tt.args.j)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Zip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZipInt(t *testing.T) {
	testZipBy(t, []zipCases[int, int, int]{
		{
			name: "base",
			args: zipArgs[int, int, int]{
				i: iter.NewEntryIterator([]iter.Entry[int, int]{
					{
						K: 0,
						V: 1,
					},
					{
						K: 0,
						V: 2,
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
				}),
			},
			want: []iter.Entry[int, iter.Entry[int, int]]{
				{K: 0, V: iter.Entry[int, int]{K: 1, V: 3}},
				{K: 1, V: iter.Entry[int, int]{K: 2, V: 1}},
			},
		},
	})
}

func TestZipPointer(t *testing.T) {
	type test struct {
		value int
	}

	testZipBy(t, []zipCases[int, *test, *test]{
		{
			name: "base",
			args: zipArgs[int, *test, *test]{
				i: iter.NewEntryIterator([]iter.Entry[int, *test]{
					{
						K: 0,
						V: &test{
							value: 1,
						},
					},
					{
						K: 0,
						V: &test{
							value: 2,
						},
					},
				}),
				j: iter.NewEntryIterator([]iter.Entry[int, *test]{
					{
						K: 0,
						V: &test{
							value: 3,
						},
					},
					{
						K: 0,
						V: &test{
							value: 4,
						},
					},
				}),
			},
			want: []iter.Entry[int, iter.Entry[*test, *test]]{
				{K: 0, V: iter.Entry[*test, *test]{K: &test{value: 1}, V: &test{value: 3}}},
				{K: 1, V: iter.Entry[*test, *test]{K: &test{value: 2}, V: &test{value: 4}}},
			},
		},
	})
}
