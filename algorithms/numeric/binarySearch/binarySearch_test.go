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
		t.Error("BinarySearch does not work properly")
	}
}
