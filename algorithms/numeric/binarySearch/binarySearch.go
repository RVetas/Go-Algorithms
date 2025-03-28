package binarySearch

// Binary Search is a search algorithm that finds the position of a target value in a sorted array.
func BinarySearch(array []int, target int) int {
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
