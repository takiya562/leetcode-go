package leetcode

import "math"

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		n := int(math.Abs(float64(num)))
		if nums[n-1] < 0 {
			res = append(res, n)
		} else {
			nums[n-1] = -nums[n-1]
		}
	}
	return res
}

func FindDuplicates(nums []int) []int {
	return findDuplicates(nums)
}
