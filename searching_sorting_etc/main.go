package main

import "fmt"

//use if condition to immpliment binary search this will resuce the time complexy of the search and will make your program faster
//take a single while loop to iterrate through the array

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid

		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(binarySearch(arr, 8), "position")
}
