package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	st := strconv.Itoa(x)
	idx := len(st) - 1
	for _, j := range st {
		if string(j) != string(st[idx]) {
			return false
		}
		idx--
	}

	return true
}

func main() {
	fmt.Println("123: \t", isPalindrome(123))
	fmt.Println("121: \t", isPalindrome(121))
	fmt.Println("12321: \t", isPalindrome(12321))
}
