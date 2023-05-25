package arrayMethods

func Pop[T any](slice *[]T) T {
	length := len(*slice)

	if length == 0 {
		var nothing T
		return nothing
	}

	removed := (*slice)[length-1]
	*slice = (*slice)[:length-1]

	return removed
}

func Shift[T any](slice *[]T) T {
	length := len(*slice)

	if length == 0 {
		var nothing T
		return nothing
	}

	removed := (*slice)[0]
	*slice = (*slice)[1:]

	return removed
}

func Reverse[T any](slice []T) {
	midIndex := (len(slice) / 2)
	var temp T
	low := 0
	high := len(slice) - 1

	for i := 0; i < midIndex; i++ {
		temp = (slice)[low]
		(slice)[low] = (slice)[high]
		(slice)[high] = temp

		low++
		high--
	}
}
