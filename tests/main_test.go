package tests

import (
	"fmt"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	a := fn.SliceIter([][]int{{1, 2, 3}, {4, 5, 6}})
	b := fn.Flatten(a)
	fmt.Println(fn.ToSlice(b))

	a1 := fn.SliceIter([][]int{{1, 2, 3}, {4, 5, 6}})
	c := fn.Map(a1, func(inK int, inV []int) (int, int) {
		return inK, len(inV)
	})
	fmt.Println(fn.ToSlice(c))
}

func TestFlatMap(t *testing.T) {
	a := fn.SliceIter([][]int{{1, 2, 3}, {4, 5, 6}})
	b := fn.FlatMap(a, func(k int, v []int) (int, []int) {
		return k, append(v, 99)
	})
	fmt.Println(fn.ToSlice(b))
}

func TestSequence(t *testing.T) {

	s := fn.SliceSeq(
		[]int{1, 2, 3},
	).Map(func(k, v int) (int, int) {
		return k, v + 1
	}).Filter(func(k, v int) bool {
		return v >= 3
	}).ToSlice()
	fmt.Println(s)

}

func TestGroupByKey(t *testing.T) {
	a := fn.SliceIter([]int{1, 2, 3, 4, 5})
	b := fn.Map[int, int](a, func(inK int, inV int) (int, int) {
		return inK % 2, inV
	})
	c := fn.GroupByKey(b)
	fmt.Println(fn.ToSlice(c))
}

func TestRange(t *testing.T) {
	fn.RangeSeq(0, 10, 1).Map(func(k, v int) (int, int) {
		return k, v + 1
	}).Filter(func(k, v int) bool { return v >= 3 }).ForEach(func(i, v int) {
		fmt.Printf("%d\n", v)
	})
}

func TestReduce(t *testing.T) {
	got := fn.RangeSeq(0, 10, 1).Reduce(func(a, b int) int {
		return a + b
	})

	assert.Equal(t, 45, got)
}
