package main

import (
	"math/rand"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
	return nums
}

func lastNumsBySlice(origin []int) []int {
	return origin[len(origin)-2:]
}

func lastNumsByCopy(origin []int) []int {
	ans := make([]int, 2)
	copy(ans, origin[len(origin)-2:])
	return ans
}
