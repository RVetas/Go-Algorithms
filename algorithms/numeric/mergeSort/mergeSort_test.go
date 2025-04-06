package mergeSort

import (
	"slices"
	"testing"
)

func TestMergeSort(t *testing.T) {
	// when
	result := MergeSort(intArray)

	// then
	if !slices.Equal(result, expectedResult) {
		t.Errorf("expected: %v, got: %v", expectedResult, result)
	}
}

func TestMergeSortStr(t *testing.T) {
	// when
	result := MergeSort(strArray)

	// then
	if !slices.Equal(result, expectedStrResult) {
		t.Errorf("expected: %v, got: %v", expectedStrResult, result)
	}
}

var (
	intArray       []int    = []int{3, 1, 7, 2, 8, 2, 6, 9, 5, 3, 1}
	expectedResult []int    = []int{1, 1, 2, 2, 3, 3, 5, 6, 7, 8, 9}
	strArray       []string = []string{
		"a8Fh2L",
		"Xv93mz",
		"pQ1nB7",
		"kLt0xR",
		"ZW3bNq",
		"uC4mY9",
		"nv5Kja",
		"Tg8wXz",
		"Jr2dML",
		"bY7qzp",
		"Vm1Eak",
		"zX0uNi",
		"cL93Pw",
		"HoX2vt",
		"qe8KrY",
		"MN7xop",
		"dVz3QW",
		"al9FtB",
		"sPw1rT",
		"Yk30xZ",
	}
	expectedStrResult []string = []string{
		"HoX2vt",
		"Jr2dML",
		"MN7xop",
		"Tg8wXz",
		"Vm1Eak",
		"Xv93mz",
		"Yk30xZ",
		"ZW3bNq",
		"a8Fh2L",
		"al9FtB",
		"bY7qzp",
		"cL93Pw",
		"dVz3QW",
		"kLt0xR",
		"nv5Kja",
		"pQ1nB7",
		"qe8KrY",
		"sPw1rT",
		"uC4mY9",
		"zX0uNi",
	}
)
