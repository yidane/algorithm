package problems

import "math"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	m := 2
	for i := 2; i <= x; i++ {
		j := x / i

		switch {
		case j < i:
			return m
		case j == i:
			return i
		case j > i:
			m = i
		}
	}

	return -1
}

//todo nice
func mySqrt1(x int) int {
	z := 1.0
	prevZ := z
	equalityThreshold := 1e-9
	firstTime := true
	for true {
		z -= ((z * z) - float64(x)) / (2 * z)
		if firstTime {
			prevZ = z
			firstTime = false
		} else {
			curZ := prevZ - z
			prevZ = z
			if math.Abs(curZ) <= equalityThreshold {
				return int(z)
			}
		}
	}
	return int(z)
}
