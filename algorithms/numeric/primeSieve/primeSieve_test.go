package primeSieve

import (
	"slices"
	"testing"
)

func TestPrimeSieve(t *testing.T) {
	// given
	expectedResult := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	// when
	result := PrimeSieve(30)
	// then
	if !slices.Equal(result, expectedResult) {
		t.Error("PrimeSieve is giving the wrong answer")
	}
}

func TestPrimeSieveSmallN(t *testing.T) {
	// given
	expectedResult := []int{}
	// when
	result := PrimeSieve(1)
	// then
	if !slices.Equal(result, expectedResult) {
		t.Error("PrimeSieve is giving the wrong answer")
	}
}
