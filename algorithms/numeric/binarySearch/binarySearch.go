package binarySearch

import "cmp"

// Binary Search is a search algorithm that finds the position of a target value in a sorted array.
func BinarySearch[T cmp.Ordered](array []T, target T) int {
	var idx int
	l := 0
	r := len(array)

	for l <= r {
		idx = (l + r) / 2
		if array[idx] == target {
			break
		}
		if array[idx] > target {
			r = idx - 1
		} else {
			l = idx + 1
		}
	}

	return idx
}

// Binary Search is a search algorithm that finds the position of a target value in a sorted array.
// This function accepts any type of arguments, but a user should provide compare function.
// Compare function returns:
// 0 if two arguments are equal
// +1 if the left argument is greater
// -1 if the right argument is greater
func BinarySearchFunc[T any](array []T, target T, compare func(lhs, rhs T) int) int {
	var idx int
	l := 0
	r := len(array)

	for l <= r {
		idx = (l + r) / 2
		if compare(array[idx], target) == 0 {
			break
		}

		if compare(array[idx], target) == 1 {
			r = idx - 1
		} else {
			l = idx + 1
		}
	}

	return idx
}
