package leetcode

func diStringMatch(s string) (ans []int) {
	l, r := 0, len(s)
	for _, c := range s {
		if c == 'D' {
			ans = append(ans, r)
			r -= 1
		}
		if c == 'I' {
			ans = append(ans, l)
			l++
		}
	}
	ans = append(ans, l)
	return
}

func DiStringMatch(s string) []int {
	return diStringMatch(s)
}
