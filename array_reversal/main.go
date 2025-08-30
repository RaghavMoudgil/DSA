package main

import "fmt"

func main() {
	arr := []int{1, 5, 15, 48, 35, 9}
	for i := 0; i < len(arr)/2; i++ {
		temp := arr[i]
		arr[i] = arr[len(arr)-1-i]
		arr[len(arr)-1-i] = temp
	}
	fmt.Println(arr)
	//
}
