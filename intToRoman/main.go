package main

import (
	"fmt"
	"strconv"
)

func intToRoman(num int) string {
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	largest := 0
	for i, v := range romanValues {
		if (num) == v {
			return i
		}

	}
	for i := 0; i < len(strconv.Itoa(num)); i++ {

	}

	return "cant convert"
}
func main() {
	fmt.Println(intToRoman(50))
}
