package main

import "fmt"

func maxSubArray(nums []int) (int, []int) {
	// subArr := make([]int, 0)
	var subArr []int
	start, end, s := 0, 0, 0
	maxSoFar, maxEndingHere := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxEndingHere+nums[i] {
			maxEndingHere = nums[i]
			s = i
		} else {
			maxEndingHere = maxEndingHere + nums[i]
		}
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
			start = s
			end = i
		}
	}
	subArr = nums[start : end+1]
	return maxSoFar, subArr
}

func main() {
	arr := []int{1, -2, 5, -3, 4, -1, 6}
	fmt.Println(maxSubArray(arr))
}
