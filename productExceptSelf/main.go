package main

import (
	"log"
)

func main() {
	log.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	suff := make([]int, n)
	pre[0] = 1
	suff[n-1] = 1

	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		suff[i] = suff[i+1] * nums[i+1]
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = pre[i] * suff[i]
	}
	return ans
}
