package primeSieve

// Sieve of Erathosthenes is a classic algorithm that allows to compute all the primes up to some integer N.
func PrimeSieve(n int) []int {
	if n < 2 {
		return []int{}
	}

	// Create initial list of numbers
	primes := make([]bool, n)
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	for i := 2; i*i < len(primes); i++ {
		if primes[i] {
			for j := i * i; j < len(primes); j += i {
				primes[j] = false
			}
		}
	}
	result := make([]int, 0, len(primes))
	for idx, val := range primes {
		if val {
			result = append(result, idx)
		}
	}
	return result
}
