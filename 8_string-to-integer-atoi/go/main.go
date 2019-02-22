package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	var sign int = 1
	var result int = 0
	var start bool = false

	for _, v := range str {
		if start == false {
			if v == '+' {
				sign = 1
				start = true
				continue
			}
			if v == '-' {
				sign = -1
				start = true
				continue
			}
		}
		if v >= '0' && v <= '9' {
			start = true
			result = result*10 + int(v-'0')
			switch {
			case sign == 1 && result > math.MaxInt32:
				return math.MaxInt32
			case sign == -1 && sign*result < math.MinInt32:
				return math.MinInt32
			}
			continue
		}
		if v != ' ' {
			break
		}
		if start == true {
			break
		}
	}

	return result * sign
}

func main() {
	fmt.Println(myAtoi("3.14"))
}
