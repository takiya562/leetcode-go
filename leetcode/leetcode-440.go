package leetcode

import "math"

func findKthNumber(n int, k int) int {
	pos, num := 1, 1
	for pos < k {
		count := childCount(n, num)
		if pos+count > k {
			num *= 10
			pos++
		} else {
			num++
			pos += count
		}
	}
	return num
}

func childCount(n, num int) int {
	cur, next := num, num+1
	count := 1
	for cur*10 <= n {
		cur *= 10
		next *= 10
		count += int(math.Min(float64(next), float64(n)+1)) - cur
	}
	return count
}

func FindKthNumber(n int, k int) int {
	return findKthNumber(n, k)
}
