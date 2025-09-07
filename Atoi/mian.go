package main

import (
	"fmt"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	var isNeg bool
	if s[0:1] == "-" {
		isNeg = true
	}
	if len(s) == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < len(s); i++ {
		if isNeg && i == 0 {
			i = i + 1
		}
		res, err := strconv.Atoi(s[i : i+1])
		if err != nil {
			if isNeg {
				ret = -ret
			}
			return ret
		}
		ret = ret*10 + res
	}
	if isNeg {
		ret = -ret
	}
	return ret
}

func main() {
	fmt.Println(myAtoi("1337c0d3"))
}
