package iter

import "errors"

type (
	Entry[T1, T2 any] struct {
		K T1
		V T2
	}

	Iterator[T1, T2 any] interface {
		Next() (Entry[T1, T2], error)
	}
)

var StopIteration = errors.New("")
var ContinueIteration = errors.New("")

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}

type Hashable interface {
	Int | Uint | ~string
}
