package iter

type MapIterator[K Comparable, V any] struct {
	Data map[K]V
	Keys []K
	Idx  int
}

func NewMapIterator[K Comparable, V any](m map[K]V) *MapIterator[K, V] {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return &MapIterator[K, V]{
		Data: m,
		Keys: keys,
		Idx:  0,
	}
}

func (i *MapIterator[K, V]) Next() (v Entry[K, V], flag Flag) {

	if i.Idx < len(i.Data) {
		v = Entry[K, V]{
			K: i.Keys[i.Idx],
			V: i.Data[i.Keys[i.Idx]],
		}
		i.Idx++
	} else {
		flag = FlagStop
	}
	return
}
