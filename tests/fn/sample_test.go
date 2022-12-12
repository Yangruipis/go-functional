package fn_test

import (
	"github.com/Yangruipis/go-functional/pkg/fn"
	"github.com/Yangruipis/go-functional/pkg/iter"

	"testing"
)

func TestSampleBase(t *testing.T) {
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
	got := fn.Entries(fn.Sample[int, int](seq, 1))
	if len(got) != 1 {
		t.Error("sample size not equal")
	}
	if got[0].V != 1 && got[0].V != 2 {
		t.Error("sample value wrong")
	}
}

func TestSampleSameSize(t *testing.T) {
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
	got := fn.Entries(fn.Sample[int, int](seq, 2))
	if len(got) != 2 {
		t.Error("sample size not equal")
	}
	if !(got[0].V == 1 && got[1].V == 2) && !(got[0].V == 2 && got[1].V == 1) {
		t.Error("sample value wrong")
	}
}

func TestSampleOutSize(t *testing.T) {
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
	got := fn.Entries(fn.Sample[int, int](seq, 3))
	if len(got) != 2 {
		t.Error("sample size not equal")
	}
	if !(got[0].V == 1 && got[1].V == 2) && !(got[0].V == 2 && got[1].V == 1) {
		t.Error("sample value wrong")
	}
}

func TestSampleProportion(t *testing.T) {
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
	got := fn.Entries(fn.Sample[int, int](seq, 0.8))
	if len(got) != 1 {
		t.Error("sample size not equal")
		return
	}
	if got[0].V != 1 && got[0].V != 2 {
		t.Error("sample value wrong")
	}
}

func TestSampleProportion2(t *testing.T) {
	seq := fn.Range(0, 100, 1)
	got := fn.Entries(fn.Sample(seq, 0.8))
	if len(got) != 80 {
		t.Error("sample size not equal")
		return
	}
}
