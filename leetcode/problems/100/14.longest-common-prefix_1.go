package problems

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	var rs []rune
	i := 0
	var si rune = -1
l:
	for {
		for _, str := range strs {
			if len(str) == 0 {
				return ""
			}

			if len(str) == i {
				break l
			}

			if si == -1 {
				si = rune(str[i])
			} else {
				if si != rune(str[i]) {
					break l
				}
			}
		}

		rs = append(rs, si)
		si = -1
		i++
	}

	return string(rs)
}
