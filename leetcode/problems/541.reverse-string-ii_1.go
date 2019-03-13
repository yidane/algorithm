package problems

func reverseStr(s string, k int) string {
	if len(s) <= 1 || k <= 1 {
		return s
	}

	rs := []byte(s)

	for i := 0; i < len(rs); i = i + 2*k {
		m, n := i, i+k-1
		if i+k > len(rs) {
			n = len(rs) - 1
		}

		for m < n {
			rs[m], rs[n] = rs[n], rs[m]
			m++
			n--
		}
	}

	return string(rs)
}
