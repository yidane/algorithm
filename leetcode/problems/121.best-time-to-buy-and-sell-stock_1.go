package problems

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (46.81%)
 * Total Accepted:    475K
 * Total Submissions: 1M
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array for which the i^th element is the price of a given
 * stock on day i.
 *
 * If you were only permitted to complete at most one transaction (i.e., buy
 * one and sell one share of the stock), design an algorithm to find the
 * maximum profit.
 *
 * Note that you cannot sell a stock before you buy one.
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Not 7-1 = 6, as selling price needs to be larger than buying price.
 *
 *
 * Example 2:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 */
//TODO worst
func maxProfit121(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	m := make(map[int]int, len(prices))
	m[0] = 0

	for i := 1; i < len(prices); i++ {
		for p1, max := range m {
			t := prices[i] - prices[p1]
			if t > max {
				m[p1] = t
			}
		}

		m[i] = 0
	}

	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}

	return max
}
