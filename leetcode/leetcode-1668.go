package leetcode

func numberOfMatches(n int) int {
	res := 0
	for n > 1 {
		match := n / 2
		res += match
		n = match + n%2
	}
	return res
}
