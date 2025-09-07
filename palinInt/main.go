package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	og := x
	rev := 0
	for x > 0 {
		remainder := x % 10
		rev = rev*10 + remainder
		x = x / 10
	}
	return rev == og
}

func main() {
	fmt.Println(isPalindrome(121))
}
