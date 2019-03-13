package problems

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 *
 * https://leetcode.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (41.69%)
 * Total Accepted:    84.1K
 * Total Submissions: 201.8K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * Given an array of integers and an integer k, you need to find the total
 * number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 *
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 *
 *
 *
 * Note:
 *
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the
 * integer k is [-1e7, 1e7].
 *
 *
 *
 */

/*
	S(i)表示第0到第i位置的子数组之和，那么S(i,j)=位置i到位置j子数组的和=S(j)-S(i-1)。
	所以，遍历一次数组，当前位置为i，那么得到了S(i)，就把S(i)存入map，map的key为和，value为和出现的次数。
	那么存入了S(i)以前，检测一下前面的数组中是否有s(i)-k存在，如果有，那么意味着i+1到j就是一个合法子数组。
	我们只需要看有没有，如果有，那么value就是这种这两个和只差的子数组的数目。最后再把s(i)存入。

	S(i)
	S(i,j)=S(j)-S(i-1)
*/

//todo 没理解
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	r, sum := 0, 0
	m := map[int]int{0: 1}

	for _, v := range nums {
		sum += v
		if c, exist := m[sum-k]; exist {
			r += c
		}
		m[sum] = m[sum] + 1
	}

	return r
}
