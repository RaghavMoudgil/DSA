package main

import "fmt"

func searchInsert(nums []int, target int) int {
	if target == nums[len(nums)-1] {
		return len(nums) - 1
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == target {
			return i
		}
	}
	for i := 0; i < len(nums)-1; i++ {

		if nums[i] != target {
			if target > nums[i] && target < nums[i+1] {
				return i + 1
			}
		}
	}
	return 0
}
func main() {
	fmt.Println(searchInsert([]int{1, 3}, 2))
}
