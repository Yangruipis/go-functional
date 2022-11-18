package iter

type (
	Entry[T1, T2 any] struct {
		K T1
		V T2
	}
	Flag int

	Iterator[T1, T2 any] interface {
		Next() (Entry[T1, T2], Flag)
	}
)

var FlagOK Flag = 0
var FlagStop Flag = 1
var FlagSkip Flag = 2

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Numeric interface {
	~float32 | ~float64
}

type Hashable interface {
	Int | Uint | Numeric | ~string
}
