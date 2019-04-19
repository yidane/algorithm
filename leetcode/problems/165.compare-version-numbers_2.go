package problems

import "strings"

/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 *
 * https://leetcode.com/problems/compare-version-numbers/description/
 *
 * algorithms
 * Medium (23.18%)
 * Total Accepted:    132.6K
 * Total Submissions: 571.9K
 * Testcase Example:  '"0.1"\n"1.1"'
 *
 * Compare two version numbers version1 and version2.
 * If version1 > version2 return 1; if version1 < version2 return -1;otherwise
 * return 0.
 *
 * You may assume that the version strings are non-empty and contain only
 * digits and the . character.
 * The . character does not represent a decimal point and is used to separate
 * number sequences.
 * For instance, 2.5 is not "two and a half" or "half way to version three", it
 * is the fifth second-level revision of the second first-level revision.
 * You may assume the default revision number for each level of a version
 * number to be 0. For example, version number 3.4 has a revision number of 3
 * and 4 for its first and second level revision number. Its third and fourth
 * level revision number are both 0.
 *
 *
 *
 * Example 1:
 *
 * Input: version1 = "0.1", version2 = "1.1"
 * Output: -1
 *
 * Example 2:
 *
 * Input: version1 = "1.0.1", version2 = "1"
 * Output: 1
 *
 * Example 3:
 *
 * Input: version1 = "7.5.2.4", version2 = "7.5.3"
 * Output: -1
 *
 * Example 4:
 *
 * Input: version1 = "1.01", version2 = "1.001"
 * Output: 0
 * Explanation: Ignoring leading zeroes, both “01” and “001" represent the same
 * number “1”
 *
 * Example 5:
 *
 * Input: version1 = "1.0", version2 = "1.0.0"
 * Output: 0
 * Explanation: The first version number does not have a third level revision
 * number, which means its third level revision number is default to "0"
 *
 *
 *
 * Note:
 *
 * Version strings are composed of numeric strings separated by dots . and this
 * numeric strings may have leading zeroes.
 * Version strings do not start or end with dots, and they will not be two
 * consecutive dots.
 *
 */
func compareVersion(version1 string, version2 string) int {
	if len(version1) == 0 {
		if len(version2) == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		if len(version2) == 0 {
			return 1
		}
	}

	vers1 := strings.Split(version1, ".")
	vers2 := strings.Split(version2, ".")

	i := 0
	for i < len(vers1) && i < len(vers2) {
		ver1 := verToRunes(vers1[i])
		ver2 := verToRunes(vers2[i])

		r := compareSubVersion(ver1, ver2)
		if r == 0 {
			i++
			continue
		}
		return r
	}

	for k := i; k < len(vers1); k++ {
		ver := verToRunes(vers1[k])
		if len(ver) != 0 {
			return 1
		}
	}

	for k := i; k < len(vers2); k++ {
		ver := verToRunes(vers2[k])
		if len(ver) != 0 {
			return -1
		}
	}

	return 0
}

func verToRunes(ver string) []rune {
	result := make([]rune, 0)
	for _, r := range ver {
		if r == '0' && len(result) == 0 {
			continue
		}

		result = append(result, r)
	}

	return result
}

func compareSubVersion(r1, r2 []rune) int {
	l1, l2 := len(r1), len(r2)

	if l1 == 0 {
		if l2 == 0 {
			return 0
		}

		return -1
	} else {
		if l2 == 0 {
			return 1
		}
	}

	if l1 > l2 {
		return 1
	}
	if l1 < l2 {
		return -1
	}

	l := 0
	if l1 > l2 {
		l = l2
	} else {
		l = l1
	}

	for i := 0; i < l; i++ {
		if r1[i] == r2[i] {
			continue
		}

		if r1[i] > r2[i] {
			return 1
		} else {
			return -1
		}
	}

	return 0
}
