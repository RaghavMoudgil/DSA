package main

import (
	"fmt"
	"math"
)

func secondLargest(num []int) int {

	first, sec, thi := math.MinInt, math.MinInt, math.MinInt

	for _, v := range num {

		if v > first {
			thi = sec
			sec = first
			first = v
		} else if v > sec && first != v {
			thi = sec
			sec = v
		} else if v > thi && sec != v {
			thi = v
		}

	}
	fmt.Println(first, sec, thi)
	return thi
}

func main() {
	arr := []int{1, 6, 5, 9, 4, 3, 894, 78, 79, 54, 1, 15}
	fmt.Print(secondLargest(arr))
}
