package binarySearch

func BinarySearch[T any](array *[]T, value T, compare func(a, b T) int) int {
	low := 0
	high := len((*array)) - 1
	mid := (low + high) / 2
	var comparison int

	for low <= high {
		comparison = compare(value, (*array)[mid])

		if comparison == 1 {
			low = mid + 1
		} else if comparison == -1 {
			high = mid - 1
		} else {
			return mid
		}

		mid = (low + high) / 2
	}

	return -1
}
