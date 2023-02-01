// Created by Bob at 2023/02/01 16:55
// https://leetcode.com/problems/valid-palindrome/

/*
125. Valid Palindrome (Easy)

A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and
removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric
characters include letters and numbers.
Given a string `s`, return `true` if it is a **palindrome**, or  `false` otherwise.
**Example 1:**
```
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```
**Example 2:**
```
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```
**Example 3:**
```
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```
**Constraints:**
- `1 <= s.length <= 2 * 10⁵`
- `s` consists only of printable ASCII characters.
*/

package main

import (
	"strings"
	"unicode"
)

// @lc code=begin

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}

		if !isAlphanumeric(s[j]) {
			j--
			continue
		}

		if !strings.EqualFold(string(s[i]), string(s[j])) {
			return false
		}
		i++
		j--
	}

	return true
}

func isAlphanumeric(b byte) bool {
	r := rune(b)
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

// @lc code=end
