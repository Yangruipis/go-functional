package iter

type SliceIterator[K int, V any] struct {
	Data []V
	Idx  int
}

func NewSliceIterator[V any](arr []V) *SliceIterator[int, V] {
	return &SliceIterator[int, V]{
		Data: arr,
		Idx:  0,
	}
}

func (i *SliceIterator[K, V]) Next() (v Entry[K, V], flag Flag) {
	if i.Idx < len(i.Data) {
		v = Entry[K, V]{
			K: K(i.Idx),
			V: i.Data[i.Idx],
		}
		i.Idx++
	} else {
		flag = FlagStop
	}
	return
}
