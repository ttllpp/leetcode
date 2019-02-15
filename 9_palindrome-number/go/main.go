package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	y := 0
	tempX := x
	for tempX > 0 {
		y = y*10 + tempX%10
		tempX /= 10
	}
	return x == y
}
