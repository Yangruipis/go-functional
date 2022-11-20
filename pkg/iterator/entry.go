package iter

type EntryIterator[T1, T2 any] struct {
	Data []Entry[T1, T2]
	Idx  int
}

func NewEntryIterator[T1, T2 any](entries []Entry[T1, T2]) *EntryIterator[T1, T2] {
	return &EntryIterator[T1, T2]{
		Data: entries,
		Idx:  0,
	}
}

func (i *EntryIterator[T1, T2]) Next() (v Entry[T1, T2], flag Flag) {
	if i.Idx < len(i.Data) {
		v = Entry[T1, T2]{
			K: i.Data[i.Idx].K,
			V: i.Data[i.Idx].V,
		}
		i.Idx++
	} else {
		flag = FlagStop
	}
	return
}
