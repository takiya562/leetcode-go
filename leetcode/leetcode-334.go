package leetcode

import "math"

func IncreasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt32
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}
	return false
}
