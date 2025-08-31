package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	start := 0
	maxlen := 0
	for idx, val := range s {
		if lastIdx, found := m[val]; found && lastIdx >= start {
			start = lastIdx + 1
		}
		m[val] = idx
		if idx-start+1 > maxlen {
			maxlen = idx - start + 1
		}
	}
	return maxlen
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
