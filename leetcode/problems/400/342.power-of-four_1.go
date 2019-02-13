package problems

func isPowerOfFour(num int) bool {
	if num == 1 || num == 4 {
		return true
	}

	b := 4

	for b < num {
		b = b * 4
	}

	return b == num
}
