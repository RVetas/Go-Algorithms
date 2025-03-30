package fibonacci

import "sync"

// Fibonacci func calculates N-th Fibonacci number.
// To speed up calculations, it uses concurrent computing and memoisation.
// Memoisation gives incredible increase in calculation speed. For N=30 concurrent implementation
// takes 1.2-1.3s to calculate (Mac M1 Pro, 32GB RAM). Concurrent implementation with memo takes 500ns.
func Fibonacci(num int64) int64 {
	if num <= 1 {
		return num
	}

	lc := make(chan int64)
	rc := make(chan int64)

	go fib(num-1, lc)
	go fib(num-2, rc)

	lv := <-lc
	rv := <-rc

	return lv + rv
}

var (
	memo  map[int64]int64 = make(map[int64]int64)
	mutex sync.Mutex
)

func fib(num int64, ch chan int64) {
	if num <= 1 {
		ch <- num
		return
	}

	// It is important to NOT lock the channels
	mutex.Lock()
	if val, ok := memo[num]; ok {
		mutex.Unlock()
		ch <- val
		return
	}
	mutex.Unlock()

	lc := make(chan int64)
	rc := make(chan int64)
	go fib(num-1, lc)
	go fib(num-2, rc)

	// just to evade ch <- (<-ch + <- ch)
	v1 := <-lc
	v2 := <-rc
	ch <- (v1 + v2)
	mutex.Lock()
	memo[num] = v1 + v2
	mutex.Unlock()
}
