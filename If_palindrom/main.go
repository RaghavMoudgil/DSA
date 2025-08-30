package main

import "fmt"

func ifPalin(str string) bool {
	run := []rune(str)

	for i, j := 0, len(run)-1; i < j; i, j = i+1, j-1 {
		run[i], run[j] = run[j], run[i]
	}
	rev := string(run)
	if rev == str {
		return true
	} else {
		return false
	}
}
func main() {
	str := "kanak"
	fmt.Println(ifPalin(str))
}
