// Created by Bob at 2023/02/01 16:42
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

/*
121. Best Time to Buy and Sell Stock (Easy)

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.
You want to maximize your profit by choosing a **single day** to buy one stock and choosing a
**different day in the future** to sell that stock.
Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit,
return `0`.
**Example 1:**
```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```
**Example 2:**
```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```
**Constraints:**
- `1 <= prices.length <= 10⁵`
- `0 <= prices[i] <= 10⁴`
*/

package main

// @lc code=begin

func maxProfit(prices []int) (ans int) {
	maxDifference := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxDifference {
			maxDifference = prices[i] - minPrice
		}
	}

	return maxDifference
}

// @lc code=end
