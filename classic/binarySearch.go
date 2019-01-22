package classic

func BinarySearch(arr []int, e int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		m := (right + left) / 2
		me := arr[m]
		switch {
		case me == e:
			return m
		case me < e:
			left = m + 1
		case me > e:
			right = m - 1
		}
	}

	return -1
}
