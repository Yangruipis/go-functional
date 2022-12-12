package fn_test

import (
	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"

	"testing"
)

func TestHead0(t *testing.T) {
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
	got := fn.Entries(fn.Head[int, int](seq, 1))
	if len(got) != 1 {
		t.Error("head size wrong")
	}
	if got[0].V != 1 {
		t.Error("head value wrong")
	}
}

func TestHead1(t *testing.T) {
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
	got := fn.Entries(fn.Head[int, int](seq, 0))
	if len(got) != 0 {
		t.Error("head size wrong")
	}
}

func TestHead2(t *testing.T) {
	seq := iter.NewEntryIterator([]iter.Entry[int, int]{})
	got := fn.Entries(fn.Head[int, int](seq, 10))
	if len(got) != 0 {
		t.Error("head size wrong")
	}
}

func TestHead3(t *testing.T) {
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
	got := fn.Entries(fn.Head[int, int](seq, 3))
	if len(got) != 2 {
		t.Error("head size wrong")
	}
	if got[0].V != 1 || got[1].V != 2 {
		t.Error("head value wrong")
	}

}

func TestHead4(t *testing.T) {
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
	got := fn.Entries(fn.Head[int, int](seq, -1))
	if len(got) != 0 {
		t.Error("head size wrong")
	}
}

func TestTail0(t *testing.T) {
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
	got := fn.Entries(fn.Tail[int, int](seq, 1))
	if len(got) != 1 {
		t.Error("head size wrong")
	}
	if got[0].V != 2 {
		t.Error("head value wrong")
	}
}

func TestTail1(t *testing.T) {
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
	got := fn.Entries(fn.Tail[int, int](seq, 2))
	if len(got) != 2 {
		t.Error("head size wrong")
	}
	if got[0].V != 2 || got[1].V != 1 {
		t.Error("head value wrong")
	}
}
