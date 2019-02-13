package problems

//todo slow
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}

	if n == 3 {
		return 1
	}

	c := 1
	primes := []int{2}
	for i := 3; i < n; i = i + 2 {
		isPrimes := true
		for _, p := range primes {
			if p*p > i {
				break
			}
			if i%p == 0 {
				isPrimes = false
				break
			}
		}

		if isPrimes {
			primes = append(primes, i)
			c++
		}
	}

	return c
}
