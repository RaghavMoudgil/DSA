package main

import "fmt"

func findPairWithSum(arr []int, target int) (int, int, bool) {
	seen := make(map[int]bool)
	for _, v := range arr {
		find := target - v
		if seen[find] {
			return find, v, true
		}
		seen[v] = true
	}
	return 0, 0, false

}

func main() {

	arr := []int{1, 6, 65, 9, 1, 615, 31, 684, 1, 651, 1, 6, 1, 65, 3, 65, 59, 2, 10, 5}

	if a, b, found := findPairWithSum(arr, 20); found {
		fmt.Println(a, b)
	} else {
		fmt.Println("no match found")
	}

}
