package main

import (
	"fmt"
)

// we are not using range loop as order is not defined in the range loop so we are usnig the below loop and method
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""
	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]     //reomving values from the main number
			result += symbols[i] // appendign symbol of the same value to the result
		}
	}
	return result
}

func romanToInt(s string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	for i := 0; i < len(s); i++ {
		currentVal := romanValues[byte(s[i])]

		if i+1 < len(s) && currentVal < romanValues[byte(s[i+1])] {
			total -= currentVal
		} else {
			total += currentVal
		}
	}
	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
