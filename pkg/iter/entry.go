package iter

type EntryIterator[K, V any] struct {
	Data []Entry[K, V]
	Idx  int
}

func NewEntryIterator[K, V any](entries []Entry[K, V]) *EntryIterator[K, V] {
	return &EntryIterator[K, V]{
		Data: entries,
		Idx:  0,
	}
}

func (i *EntryIterator[K, V]) Next() (v Entry[K, V], flag Flag) {
	if i.Idx < len(i.Data) {
		v = Entry[K, V]{
			K: i.Data[i.Idx].K,
			V: i.Data[i.Idx].V,
		}
		i.Idx++
	} else {
		flag = FlagStop
	}
	return
}
