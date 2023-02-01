// Created by Bob at 2023/02/01 16:54
// https://leetcode.com/problems/valid-parentheses/

/*
20. Valid Parentheses (Easy)

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`,
determine if the input string is valid.
An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.
**Example 1:**
```
Input: s = "()"
Output: true
```
**Example 2:**
```
Input: s = "()[]{}"
Output: true
```
**Example 3:**
```
Input: s = "(]"
Output: false
```
**Constraints:**
- `1 <= s.length <= 10⁴`
- `s` consists of parentheses only `'()[]{}'`.
*/

package main

// @lc code=begin

func isValid(s string) bool {
	closingRune := make(map[rune]rune, 3)
	closingRune['['] = ']'
	closingRune['{'] = '}'
	closingRune['('] = ')'

	openStack := make([]rune, 0, len(s))
	for _, r := range s {
		var (
			lastOpenedRune, expectedClosingRune rune
			lenOpenStack                        = len(openStack)
		)

		if lenOpenStack > 0 {
			lastOpenedRune = openStack[lenOpenStack-1]
			expectedClosingRune = closingRune[lastOpenedRune]
		}

		if r == '[' || r == '{' || r == '(' {
			openStack = append(openStack, r)
		} else if r == expectedClosingRune {
			openStack = openStack[:lenOpenStack-1]
		} else {
			// unexpected closing parentheses
			return false
		}
	}

	return len(openStack) == 0
}

// @lc code=end
