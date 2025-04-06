package mergeSort

import (
	"cmp"
)

func MergeSort[T cmp.Ordered](array []T) []T {
	if len(array) < 2 {
		return array
	}

	result := make([]T, len(array), len(array))

	l := MergeSort(array[:len(array)/2])
	r := MergeSort(array[len(array)/2:])

	counter := 0
	for len(l) > 0 && len(r) > 0 {
		if l[0] < r[0] {
			result[counter] = l[0]
			l = l[1:]
		} else {
			result[counter] = r[0]
			r = r[1:]
		}

		counter++
	}

	for len(l) > 0 {
		result[counter] = l[0]
		l = l[1:]
		counter++
	}

	for len(r) > 0 {
		result[counter] = r[0]
		r = r[1:]
		counter++
	}
	return result
}
