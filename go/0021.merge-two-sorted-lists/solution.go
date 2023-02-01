// Created by Bob at 2023/02/01 16:53
// https://leetcode.com/problems/merge-two-sorted-lists/

/*
21. Merge Two Sorted Lists (Easy)

You are given the heads of two sorted linked lists `list1` and `list2`.
Merge the two lists in a one **sorted** list. The list should be made by splicing together the nodes
of the first two lists.
Return the head of the merged linked list.
**Example 1:**
![](https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg)
```
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
```
**Example 2:**
```
Input: list1 = [], list2 = []
Output: []
```
**Example 3:**
```
Input: list1 = [], list2 = [0]
Output: [0]
```
**Constraints:**
- The number of nodes in both lists is in the range `[0, 50]`.
- `-100 <= Node.val <= 100`
- Both `list1` and `list2` are sorted in **non-decreasing** order.
*/

package main

import . "github.com/j178/leetgo/testutils/go"

// @lc code=begin

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var list3, head *ListNode
	if list1.Val <= list2.Val {
		list3 = list1
		list1 = list1.Next
	} else {
		list3 = list2
		list2 = list2.Next
	}
	head = list3

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			list3.Next = list1
			list3 = list1
			list1 = list1.Next
		} else {
			list3.Next = list2
			list3 = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		list3.Next = list1
	}

	if list2 != nil {
		list3.Next = list2
	}

	return head
}

// @lc code=end
