Given a string s, find the length of the longest substring without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
Example 4:

Input: s = ""
Output: 0

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.

## Notes

I struggled with when to add the reset / modification steps like adding onto the list.

I feel like getting the base case of one element was hacky.

The LeetCode submission metrics said that my code was only about 10% faster than the rest.

I tripped up on not thinking about simple use cases.

The term that I keep seeing in the discussion is 'sliding window'.

A substring and a subsequence ARE NOT THE SAME THING. Needed to understand this.

Its not that easy to write nested loops of which have indices that traverse a collection in pairs.
