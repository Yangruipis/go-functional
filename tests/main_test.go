package tests

import (
	"fmt"
	"testing"

	fun "github.com/Yangruipis/go-functional/pkg"
	"github.com/stretchr/testify/assert"
)

func TestFlatten(t *testing.T) {
	a := fun.NewSlice([][]int{{1, 2, 3}, {4, 5, 6}})
	b := fun.Flatten(a)
	fmt.Println(fun.ToSlice(b))

	a1 := fun.NewSlice([][]int{{1, 2, 3}, {4, 5, 6}})
	c := fun.Map(a1, func(inK int, inV []int) (int, int) {
		return inK, len(inV)
	})
	fmt.Println(fun.ToSlice(c))
}

func TestFlatMap(t *testing.T) {
	a := fun.NewSlice([][]int{{1, 2, 3}, {4, 5, 6}})
	b := fun.FlatMap(a, func(k int, v []int) (int, []int) {
		return k, append(v, 99)
	})
	fmt.Println(fun.ToSlice(b))
}

func TestSequence(t *testing.T) {

	s := fun.SliceSeq(
		[]int{1, 2, 3},
	).Map(func(k, v int) (int, int) {
		return k, v + 1
	}).Filter(func(k, v int) bool {
		return v >= 3
	}).ToSlice()
	fmt.Println(s)

}

func TestGroupByKey(t *testing.T) {
	a := fun.NewSlice([]int{1, 2, 3, 4, 5})
	b := fun.Map[int, int](a, func(inK int, inV int) (int, int) {
		return inK % 2, inV
	})
	c := fun.GroupByKey(b)
	fmt.Println(fun.ToSlice(c))
}

func TestRange(t *testing.T) {
	fun.RangeSeq(0, 10, 1).Map(func(k, v int) (int, int) {
		return k, v + 1
	}).Filter(func(k, v int) bool { return v >= 3 }).ForEach(func(i, v int) {
		fmt.Printf("%d\n", v)
	})
}

func TestReduce(t *testing.T) {
	got := fun.RangeSeq(0, 10, 1).Reduce(func(a, b int) int {
		return a + b
	})

	assert.Equal(t, 45, got)
}
