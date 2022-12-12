package fn_test

import (
	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"

	"testing"
)

func TestChoicesBase(t *testing.T) {
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
	got := fn.Entries(fn.Choices[int, int](seq, 1))
	if len(got) != 1 {
		t.Error("choices size not equal")
	}
	if got[0].V != 1 && got[0].V != 2 {
		t.Error("choices value wrong")
	}
}

func TestChoicesSameSize(t *testing.T) {
	seq := []int{1, 2}

	hitPattens := make(map[int]bool)
	for i := 0; i < 100; i++ {
		got := fn.Entries(fn.Choices[int, int](iter.NewSliceIterator(seq), 2))
		if len(got) != 2 {
			t.Error("choices size not equal")
		}
		if got[0].V == 1 && got[1].V == 2 {
			hitPattens[0] = true
		} else if got[0].V == 1 && got[1].V == 1 {
			hitPattens[1] = true
		} else if got[0].V == 2 && got[1].V == 2 {
			hitPattens[2] = true
		} else if got[0].V == 2 && got[1].V == 1 {
			hitPattens[3] = true
		}
	}
	for i := 0; i < 4; i++ {
		ok, v := hitPattens[i]
		if ok && v {
			continue
		} else {
			t.Error("patten missed")
		}
	}
}

func TestChoicesOutSize(t *testing.T) {
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
	got := fn.Entries(fn.Choices[int, int](seq, 3))
	if len(got) != 3 {
		t.Error("choices size not equal")
	}
}

func TestChoicesProportion2(t *testing.T) {
	seq := fn.Range(0, 100, 1)
	got := fn.Entries(fn.Choices(seq, 0.8))
	if len(got) != 80 {
		t.Error("choices size not equal")
		return
	}
}
