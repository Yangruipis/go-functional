package iter

type MapIterator[T1 Hashable, T2 any] struct {
	Data map[T1]T2
	Keys []T1
	Idx  int
}

func NewMapIterator[T1 Hashable, T2 any](m map[T1]T2) *MapIterator[T1, T2] {
	keys := make([]T1, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &MapIterator[T1, T2]{
		Data: m,
		Keys: keys,
		Idx:  0,
	}
}

func (i *MapIterator[T1, T2]) Next() (v Entry[T1, T2], flag Flag) {

	if i.Idx < len(i.Data) {
		v = Entry[T1, T2]{
			K: i.Keys[i.Idx],
			V: i.Data[i.Keys[i.Idx]],
		}
		i.Idx++
	} else {
		flag = FlagStop
	}
	return
}
