package leetcode

func maxDepth(s string) (ans int) {
	max := 0
	for _, ch := range s {
		if ch == '(' {
			max++
			if max > ans {
				ans = max
			}
		} else if ch == ')' {
			max--
		}
	}
	return
}
