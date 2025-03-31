package fibonacci

import (
	"fmt"
	"testing"
	"time"
)

func TestFibonacci(t *testing.T) {
	// given
	expectedResult := 6334266236422402381
	// when
	start := time.Now()
	result := Fibonacci(103)
	// then
	fmt.Printf("calculation time: %s\n", time.Since(start))
	if result != expectedResult {
		t.Errorf("expected %d, got %d\n", expectedResult, result)
	}
}
