// Created by Bob at 2023/02/17 16:04
// https://leetcode.com/problems/maximum-subarray/

/*
53. Maximum Subarray (Medium)

Given an integer array `nums`, find the subarray with the largest sum, and return its sum.
**Example 1:**
```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.
```
**Example 2:**
```
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.
```
**Example 3:**
```
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.
```
**Constraints:**
- `1 <= nums.length <= 10⁵`
- `-10⁴ <= nums[i] <= 10⁴`
**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the
**divide and conquer** approach, which is more subtle.
*/

package main

// @lc code=begin

func maxSubArray(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}

		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}

// @lc code=end
