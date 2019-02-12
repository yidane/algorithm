package problems

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}

	as := make([]rune, 0)
	v := func(r rune) bool {
		if len(as) == 0 {
			return false
		}

		f := as[len(as)-1] == r
		as = as[:len(as)-1]

		return f
	}

	for i := 0; i < len(s); i++ {
		switch rune(s[i]) {
		case '(':
			as = append(as, ')')
		case ')':
			if !v(')') {
				return false
			}
		case '[':
			as = append(as, ']')
		case ']':
			if !v(']') {
				return false
			}
		case '{':
			as = append(as, '}')
		case '}':
			if !v('}') {
				return false
			}
		default:
			return false
		}
	}

	return len(as) == 0
}
