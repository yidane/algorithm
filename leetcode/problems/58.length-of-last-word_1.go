package problems

func lengthOfLastWord(s string) int {
	l := 0

	space := byte(' ')
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == space {
			if l != 0 {
				break
			}
		} else {
			l++
		}
	}

	return l
}
