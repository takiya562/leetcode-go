package search

func BinarySearchRight(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return end
}

func BinarySearchLeft(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func BinarySearch(nums []int, start, end, target int) int {
	if start > end {
		return -1
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid
		}
	}
	if nums[start] == target {
		return start
	}
	return -1
}
