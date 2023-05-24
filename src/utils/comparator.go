package comparator

type Comparison int

const (
	BIGGER  Comparison = 1
	EQUAL              = 0
	SMALLER            = -1
)

type CompareFn[T any] func(a T, b T) Comparison
