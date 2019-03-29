package problems

/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (31.39%)
 * Total Accepted:    396.1K
 * Total Submissions: 1.3M
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) < len(needle) {
		return -1
	}

	c1 := needle[0]
	for i := 0; i < len(haystack); i++ {
		if len(haystack)-i < len(needle) {
			return -1
		}

		if haystack[i] == c1 {
			match := true
			for j := 1; j < len(needle); j++ {
				if needle[j] != haystack[i+j] {
					match = false
					break
				}
			}

			if match {
				return i
			}
		}
	}

	return -1
}

/*
//其实这个你自己静下心模拟一下过程，可以很快理解的
void kmp_next(char *s)
{
    int len = strlen(s);
    next[0] = -1;
    int k = -1;
    for(int i = 1; i < len; i ++) {
        while(k > -1 && s[k+1] != s[i]) {
            k = next[k];
        }
        if(s[k+1] == s[i]) {
            k ++;
        }
        next[i] = k;
    }
}

public void search(String original, String find, int next[]) {
	int j = 0;
	for (int i = 0; i < original.length(); i++) {
		while (j > 0 && original.charAt(i) != find.charAt(j))
			j = next[j];
		if (original.charAt(i) == find.charAt(j))
			j++;
		if (j == find.length()) {
			System.out.println("find at position " + (i - j));
			System.out.println(original.subSequence(i - j + 1, i + 1));
			j = next[j];
		}
	}
}
*/

//TODO:KMP算法未理解
func kmpNext(s string) []int {
	next := make([]int, len(s))
	next[0] = 0
	k := 0

	for i := 1; i < len(s); i++ {
		for k > 0 && s[i] != s[k] {
			k = next[k]
		}

		if s[k] == s[i] {
			k++
		}

		next[i] = k
	}

	return next
}

func search(original, find string, next []int) int {
	j := 0

	for i := 0; i < len(original); i++ {
		for j > 0 && original[i] != find[j] {
			j = next[j]
		}

		if original[i] == find[j] {
			j++
		}

		if j == len(find) {
			return i - j
		}
	}

	return -1
}

func KMP(haystack string, needle string) int {
	next := kmpNext(needle)
	return search(haystack, needle, next)
}
