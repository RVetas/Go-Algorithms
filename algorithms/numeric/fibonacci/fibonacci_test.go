package fibonacci

import (
	"fmt"
	"testing"
	"time"
)

func TestFibonacci(t *testing.T) {
	// given
	expectedResult := int64(6334266236422402381)
	// when
	start := time.Now()
	result := Fibonacci(103)
	fmt.Printf("calculation time: %s\n", time.Since(start))
	// then
	if result != expectedResult {
		t.Errorf("expected %d, got %d", expectedResult, result)
	}
}
