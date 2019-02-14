package main

import (
	"math"
)

func reverse(x int) int {
	result := 0
	for x != 0 {
		temp := x % 10
		result = 10*result + temp
		x /= 10 //如果写成x = x / 10 内存占用大一倍，但是性能不影响，写成x = x / 10的时候,go的编译器会做优化
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}
