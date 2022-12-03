package fn_test

import (
	"reflect"
	"testing"

	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"
)

func TestShuffle(t *testing.T) {

	hit0 := false
	hit1 := false

	for i := 1; i <= 100; i++ {
		seq := iter.NewEntryIterator([]iter.Entry[int, int]{
			{
				K: 0,
				V: 1,
			},
			{
				K: 1,
				V: 2,
			},
		})
		got := fn.Entries(fn.Shuffle[int, int](seq))
		if reflect.DeepEqual(got, []iter.Entry[int, int]{
			{
				K: 0,
				V: 1,
			},
			{
				K: 1,
				V: 2,
			},
		}) {
			hit0 = true
		}
		if reflect.DeepEqual(got, []iter.Entry[int, int]{
			{
				K: 1,
				V: 2,
			},
			{
				K: 0,
				V: 1,
			},
		}) {
			hit1 = true
		}
	}
	if !hit0 || !hit1 {
		t.Error("not shuffled")
	}
}
