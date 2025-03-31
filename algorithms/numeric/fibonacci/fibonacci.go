package fibonacci

var (
	memo map[int]int = make(map[int]int)
)

// Fibonacci func computes N-th Fibonacci number. It also utilizes memoisation to speed up computations.
func Fibonacci(num int) int {
	if num < 2 {
		return num
	}

	if val, ok := memo[num]; ok {
		return val
	}

	v1 := Fibonacci(num - 1)
	v2 := Fibonacci(num - 2)
	memo[num] = v1 + v2
	return v1 + v2
}
