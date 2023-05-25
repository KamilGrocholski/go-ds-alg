package quickSort

type LessFn[T any] func(a, b T) bool

func QuickSort[T any](array *[]T, less LessFn[T]) {
	quickSortHelper(array, 0, len((*array))-1, less)
}

func quickSortHelper[T any](array *[]T, low int, high int, less LessFn[T]) {
	if low > high {
		return
	}

	pivotIndex := partition(array, low, high, less)
	quickSortHelper(array, low, pivotIndex-1, less)
	quickSortHelper(array, pivotIndex+1, high, less)
}

func partition[T any](array *[]T, low int, high int, less LessFn[T]) int {
	pivotValue := (*array)[high]
	partitionIndex := low - 1

	for i := low; i < high; i++ {
		if less((*array)[i], pivotValue) == true {
			partitionIndex++
			temp := (*array)[i]
			(*array)[i] = (*array)[partitionIndex]
			(*array)[partitionIndex] = temp
		}
	}

	partitionIndex++
	(*array)[high] = (*array)[partitionIndex]
	(*array)[partitionIndex] = pivotValue

	return partitionIndex
}
