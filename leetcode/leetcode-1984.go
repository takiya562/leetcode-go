package leetcode

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	min := math.MaxInt32
	for i := 0; i < n-k+1; i++ {
		gap := nums[i+k-1] - nums[i]
		if min > gap {
			min = gap
		}
	}
	return min
}

func MinimumDifference(nums []int, k int) int {
	return minimumDifference(nums, k)
}
