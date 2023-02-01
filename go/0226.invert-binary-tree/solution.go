// Created by Bob at 2023/02/01 16:46
// https://leetcode.com/problems/invert-binary-tree/

/*
226. Invert Binary Tree (Easy)

Given the `root` of a binary tree, invert the tree, and return its root.
**Example 1:**
![](https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg)
```
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]
```
**Example 2:**
![](https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg)
```
Input: root = [2,1,3]
Output: [2,3,1]
```
**Example 3:**
```
Input: root = []
Output: []
```
**Constraints:**
- The number of nodes in the tree is in the range `[0, 100]`.
- `-100 <= Node.val <= 100`
*/

package main

import (
	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func invertTree(root *TreeNode) (ans *TreeNode) {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}

// @lc code=end
