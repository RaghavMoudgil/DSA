package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	rev := 0
	neg := x < 0
	if neg {
		x = -x
	}
	for x > 0 {
		remainder := x % 10
		rev = rev*10 + remainder
		x = x / 10

	}
	if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10) {
		return 0
	}
	if neg {
		rev = -rev
	}
	return rev
}

func main() {
	fmt.Print(reverse(-123))
}
