package leetcode

func singleNonDuplicate(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) >> 1
		if mid%2 == 0 {
			if nums[mid] != nums[mid+1] {
				r = mid
			} else {
				l = mid + 2
			}
		} else {
			if nums[mid] != nums[mid-1] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return nums[l]
}

func SingleNonDuplicate(nums []int) int {
	return singleNonDuplicate(nums)
}
