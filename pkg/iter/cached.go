package iter

type CachedIterator[K, V any] struct {
	Data []Entry[K, V]
	Idx  int
}

func NewCachedIterator[K, V any](entries []Entry[K, V]) *CachedIterator[K, V] {
	return &CachedIterator[K, V]{
		Data: entries,
		Idx:  0,
	}
}

func (i *CachedIterator[K, V]) Next() (v Entry[K, V], flag Flag) {
	if i.Idx < len(i.Data) {
		v = Entry[K, V]{
			K: i.Data[i.Idx].K,
			V: i.Data[i.Idx].V,
		}
		i.Idx++
	} else {
		flag = FlagStop
		i.Idx = 0 // make this able to reuse
	}
	return
}
