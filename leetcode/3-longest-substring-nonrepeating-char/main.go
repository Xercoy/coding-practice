package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var substrLen int

	// handle base case of single element
	if len(s) == 1 {
		substrLen = 1
	}

	for i := 0; i <= len(s)-2; i++ {
		// make a map to keep track of substring
		charList := make(map[byte]int)
		charList[s[i]] += 1

		for j := i + 1; j < len(s); j++ {
			// this only happens when you hit a duplicate
			if charList[s[j]] > 0 {
				break
			} else {
				charList[s[j]] += 1
			}
		}

		listLen := len(charList)
		if listLen > substrLen {
			substrLen = listLen
		}
	}

	return substrLen
}

func main() {
	cases := []struct {
		str      string
		expected int
	}{
		{"aa", 1},
		{"", 0},
		{"au", 2},
		{" ", 1},
		{"bbbbbb", 1},
	}

	for i, c := range cases {
		if lengthOfLongestSubstring(c.str) != c.expected {
			fmt.Printf("\ncase #%d FAIL", i)
		} else {
			fmt.Printf("\ncase #%d PASS", i)
		}
	}
}
