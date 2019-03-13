package problems

//slow
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

//good version
func countPrimes1(n int) int {
	if n < 2 {
		return 0
	}
	table := make([]bool, n+1)
	table[0], table[1] = true, true
	numPrimes := 0
	for i := 2; i < n; i++ {
		if table[i] {
			continue
		}
		numPrimes++
		for j := i * i; j < n; j += i { //质数的倍数不为质数
			table[j] = true
		}
	}
	return numPrimes
}
