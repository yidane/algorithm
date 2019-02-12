package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	if x%10 == 0 {
		return false
	}

	o := x
	n := 0

	for x >= 10 {
		i := x % 10
		n = n*10 + i
		x = (x - i) / 10
	}

	n = n*10 + x

	return o == n
}
