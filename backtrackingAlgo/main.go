package main

import "fmt"

func letterCombinations(digits string) []string {
	mapping := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}

	if digits == "" {
		return []string{}

	}
	var result []string
	var backTrack func(index int, curr string)
	backTrack = func(index int, curr string) {
		if index == len(digits) {
			result = append(result, curr)
		}
		letter := mapping[digits[index]]
		for i := 0; i < len(letter); i++ {
			backTrack(index+1, curr+string(letter[i]))
		}
	}
	backTrack(0, "")
	return result

}

func main() {
	fmt.Println(letterCombinations("23"))
}
