package leetcode

var N int = 100010
var c []int

func bestRotation(nums []int) int {
	c = make([]int, N)
	n := len(nums)
	for i, num := range nums {
		a, b := (i+1)%n, (i-num+n)%n
		if a <= b {
			add(a, b)
		} else {
			add(0, b)
			add(a, n-1)
		}
	}
	for i := 1; i < n; i++ {
		c[i] += c[i-1]
	}
	ans := 0
	for i := 1; i < N; i++ {
		if c[i] > c[ans] {
			ans = i
		}
	}
	return ans
}

func add(l, r int) {
	c[l]++
	c[r+1]--
}

func BestRotation(nums []int) int {
	return bestRotation(nums)
}
