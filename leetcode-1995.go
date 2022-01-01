package main

import (
	"fmt"
)

func countQuadruplets(nums []int) (ans int) {
	cnt := map[int]int{}
	for b := len(nums) - 3; b >= 1; b-- {
		for _, x := range nums[b+2:] {
			cnt[x-nums[b+1]]++
		}
		for _, x := range nums[:b] {
			if sum := x + nums[b]; cnt[sum] > 0 {
				ans += cnt[sum]
			}
		}
	}
	return
}

func main() {
	nums := []int{28, 8, 49, 85, 37, 90, 20, 8}
	fmt.Println(nums[:3])
}
