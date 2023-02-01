// Created by Bob at 2023/02/01 16:44
// https://leetcode.com/problems/binary-search/

/*
704. Binary Search (Easy)

Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write
a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise,
return `-1`.
You must write an algorithm with `O(log n)` runtime complexity.
**Example 1:**
```
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
```
**Example 2:**
```
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
```
**Constraints:**
- `1 <= nums.length <= 10⁴`
- `-10⁴ < nums[i], target < 10⁴`
- All the integers in `nums` are **unique**.
- `nums` is sorted in ascending order.
*/

package main

// @lc code=begin

func search(nums []int, target int) (ans int) {
	min := 0
	max := len(nums) - 1

	for min <= max {
		mid := (min + max) / 2
		if nums[mid] < target {
			min = mid + 1
		} else if nums[mid] > target {
			max = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

// @lc code=end
