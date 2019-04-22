package problems

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 *
 * https://leetcode.com/problems/valid-palindrome/description/
 *
 * algorithms
 * Easy (30.73%)
 * Total Accepted:    342.6K
 * Total Submissions: 1.1M
 * Testcase Example:  '"A man, a plan, a canal: Panama"'
 *
 * Given a string, determine if it is a palindrome, considering only
 * alphanumeric characters and ignoring cases.
 *
 * Note: For the purpose of this problem, we define empty string as valid
 * palindrome.
 *
 * Example 1:
 *
 *
 * Input: "A man, a plan, a canal: Panama"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "race a car"
 * Output: false
 *
 *
 */
func isPalindrome125(s string) bool {
	var arr []uint8

	for i := 0; i < len(s); i++ {
		si := s[i]
		if si >= '0' && si <= '9' {
			arr = append(arr, si)
			continue
		}

		if si >= 'a' && si <= 'z' {
			arr = append(arr, si)
			continue
		}

		if si >= 'A' && si <= 'Z' {
			arr = append(arr, si+32)
		}
	}

	if len(arr) == 0 {
		return true
	}

	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindrome125_1(s string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; {
		left := alphanumeric(s[i])
		if left == 0 {
			i++
			continue
		}

		right := alphanumeric(s[j])
		if right == 0 {
			j--
			continue
		}

		if left != right {
			return false
		}

		i++
		j--
	}

	return true
}

func alphanumeric(r uint8) uint8 {
	if r >= '0' && r <= '9' {
		return r
	}

	if r >= 'a' && r <= 'z' {
		return r
	}

	if r >= 'A' && r <= 'Z' {
		r += 32
		return r
	}

	return 0
}
