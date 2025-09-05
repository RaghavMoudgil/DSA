package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	first := needle[0:1]
	length := len(needle)
	for i := 0; i < len(haystack); i++ {
		temp := haystack[i : i+1]
		if temp == first {
			new := haystack[i : i+length]
			if new == needle {
				return i
			}
		}
	}

	return 0

}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))

}
