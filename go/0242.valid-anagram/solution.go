// Created by Bob at 2023/02/01 16:57
// https://leetcode.com/problems/valid-anagram/

/*
242. Valid Anagram (Easy)

Given two strings `s` and `t`, return `true`if `t`is an anagram of `s`, and `false`otherwise.
An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
**Example 1:**
```
Input: s = "anagram", t = "nagaram"
Output: true
```
**Example 2:**
```
Input: s = "rat", t = "car"
Output: false
```
**Constraints:**
- `1 <= s.length, t.length <= 5 * 10⁴`
- `s` and `t` consist of lowercase English letters.
**Follow up:** What if the inputs contain Unicode characters? How would you adapt your solution to
such a case?
*/

package main

// @lc code=begin

func isAnagram(s string, t string) bool {
	occurrencesByLetter := make(map[rune]int16)
	for _, letter := range s {
		occurrencesByLetter[letter]++
	}

	for _, letter := range t {
		occurrencesByLetter[letter]--
	}

	for _, occurrences := range occurrencesByLetter {
		if occurrences != 0 {
			return false
		}
	}

	return true
}

// @lc code=end
