package problems

/*
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
Input: "aba"
Output: True
Example 2:
Input: "abca"
Output: True
Explanation: You could delete the character 'c'.
Note:
The string will only contain lowercase characters a-z. The maximum length of the string is 50000.
*/

func validPalindrome(s string) bool {
	delCount := 0
	i, j := 0, len(s)-1
	flag := false
	for i < j {
		if s[i] != s[j] {
			delCount++
			if delCount > 1 {
				return false
			}

			if s[i+1] != s[j] && s[i] != s[j-1] {
				return false
			}

			if s[i+1] == s[j] {
				if j > i+2 && s[i+2] != s[j-1] {
					flag = true
				}

				if !flag {
					i++
					continue
				}
			}

			j--
			flag = false
		} else {
			i++
			j--
		}
	}

	return true
}
