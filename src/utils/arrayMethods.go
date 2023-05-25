package arrayMethods

type Slice[T any] []T

func (slice *Slice[T]) Pop() T {
	length := len(*slice)

	if length == 0 {
		var nothing T
		return nothing
	}

	removed := (*slice)[length-1]
	*slice = (*slice)[:length-1]

	return removed
}

func (slice *Slice[T]) Shift() T {
	length := len(*slice)

	if length == 0 {
		var nothing T
		return nothing
	}

	removed := (*slice)[0]
	*slice = (*slice)[1:]

	return removed
}
