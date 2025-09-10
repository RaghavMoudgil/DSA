package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return "" // Handle empty input array
	}
	if len(strs) == 1 {
		return strs[0] // If only one string, it's the prefix
	}

	// Sort the array of strings. This makes the first and last strings
	// in the sorted array have the longest common prefix with all other strings.
	sort.Strings(strs)

	first := strs[0]
	last := strs[len(strs)-1]
	minLength := min(len(first), len(last))

	// Iterate through characters of the first and last strings
	// to find the common prefix.
	i := 0
	for i < minLength && first[i] == last[i] {
		i++
	}

	// Return the common prefix substring.
	return first[0:i]
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
