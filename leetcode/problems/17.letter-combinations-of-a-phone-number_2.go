package problems

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (40.81%)
 * Total Accepted:    363.9K
 * Total Submissions: 890.5K
 * Testcase Example:  '"23"'
 *
 * Given a string containing digits from 2-9 inclusive, return all possible
 * letter combinations that the number could represent.
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below. Note that 1 does not map to any letters.
 *
 *
 *
 * Example:
 *
 *
 * Input: "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * Note:
 *
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	telephone := map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var result = telephone[digits[0]]

	for i := 1; i < len(digits); i++ {
		result = letterCombinationsPackage(result, telephone[digits[i]])
	}

	return result
}

func letterCombinationsPackage(arr1, arr2 []string) []string {
	result := make([]string, len(arr1)*len(arr2))

	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			result[i*len(arr2)+j] = arr1[i] + arr2[j]
		}
	}

	return result
}
