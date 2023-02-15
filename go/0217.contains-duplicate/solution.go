// Created by Bob at 2023/02/15 19:24
// https://leetcode.com/problems/contains-duplicate/

/*
217. Contains Duplicate (Easy)

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array,
and return `false` if every element is distinct.
**Example 1:**
```
Input: nums = [1,2,3,1]
Output: true
```
**Example 2:**
```
Input: nums = [1,2,3,4]
Output: false
```
**Example 3:**
```
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
```
**Constraints:**
- `1 <= nums.length <= 10⁵`
- `-10⁹ <= nums[i] <= 10⁹`
*/

package main

// @lc code=begin

func containsDuplicate(nums []int) bool {
	existingNums := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := existingNums[num]; ok {
			return true
		}
		existingNums[num] = struct{}{}
	}
	return false
}

// @lc code=end
