package binarySearch

import (
	"math/rand"
	"slices"
	"testing"
)

// Lets create a sorted array of length N (=testSize), that contains numbers from 0 to 99.
// Then we use BinarySearchFunc N (=testSize) times to try and find an index of each number from 0 to N-1,
// which (very conviniently) equalts to that number. This allows us to simply compare two arrays.
func TestBinarySearch(t *testing.T) {
	// given
	testSize := rand.Intn(10000)
	arrayToSearch := func() []int {
		array := make([]int, testSize, testSize)
		for i := 0; i < testSize; i++ {
			array[i] = i
		}
		return array
	}()

	// when
	resultsArray := make([]int, testSize, testSize)
	for i := 0; i < testSize; i++ {
		resultsArray[i] = BinarySearch(arrayToSearch, i)
	}

	// then
	if !slices.Equal(resultsArray, arrayToSearch) {
		t.Errorf("BinarySearch does not work properly\n expected: %v, got: %v\n", arrayToSearch, resultsArray)
	}
}

// This test is a bit more sophisticated version of a previos Binary Search test.
// We introduce TestType that does not conform to cmp.Ordered generic constraint and a compareFunc that
// we will use in BinarySearchFunc as a comparator.
// Otherwise, test logic is the same as in the test above.
func TestBinarySearchFunc(t *testing.T) {
	type TestType struct {
		Val int
	}

	compareFunc := func(lhs, rhs TestType) int {
		if lhs.Val == rhs.Val {
			return 0
		} else if lhs.Val > rhs.Val {
			return +1
		} else {
			return -1
		}
	}

	// given
	testSize := rand.Intn(10000)
	arrayToSearch := func() []TestType {
		array := make([]TestType, testSize, testSize)
		for i := 0; i < testSize; i++ {
			array[i] = TestType{Val: i}
		}
		return array
	}()
	// when
	resultsArray := make([]int, testSize, testSize)
	for i := 0; i < testSize; i++ {
		resultsArray[i] = BinarySearchFunc(arrayToSearch, TestType{Val: i}, compareFunc)
	}

	// then
	if !slices.EqualFunc(resultsArray, arrayToSearch, func(i int, t TestType) bool { return i == t.Val }) {
		t.Errorf("BinarySearch does not work properly\n expected: %v, got: %v\n", arrayToSearch, resultsArray)
	}
}
