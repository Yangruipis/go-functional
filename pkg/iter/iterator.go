package iter

type (
	Entry[K, V any] struct {
		K K
		V V
	}
	Flag int

	Iterator[K, V any] interface {
		Next() (Entry[K, V], Flag)
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

type Comparable interface {
	Int | Uint | Numeric | ~string
}
