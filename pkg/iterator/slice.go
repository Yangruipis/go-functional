package iter

type SliceIterator[T1 int, T2 any] struct {
	Data []T2
	Idx  int
}

func NewSliceIterator[T2 any](arr []T2) *SliceIterator[int, T2] {
	return &SliceIterator[int, T2]{
		Data: arr,
		Idx:  0,
	}
}

func (i *SliceIterator[T1, T2]) Next() (v Entry[T1, T2], err error) {
	if i.Idx < len(i.Data) {
		v = Entry[T1, T2]{
			K: T1(i.Idx),
			V: i.Data[i.Idx],
		}
		i.Idx++
	} else {
		err = StopIteration
	}
	return
}
