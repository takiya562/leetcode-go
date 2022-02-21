package leetcode

func pushDominoes(dominoes string) string {
	n := len(dominoes)
	dp := make([]int, n)
	for i, c := range dominoes {
		if c == '.' {
			dp[i] = 0
		} else if c == 'R' {
			dp[i] = 1
		} else {
			dp[i] = -1
		}
		if dp[i] == 0 && i > 0 && dp[i-1] > 0 {
			dp[i] = dp[i-1] + 1
		}
	}
	res := make([]byte, n)
	for i := n - 1; i >= 0; i-- {
		if dp[i] == 0 && i < n-1 && dp[i+1] < 0 {
			dp[i] = dp[i+1] - 1
		}
		if i < n-1 && dp[i+1] < 0 && dp[i] > 0 {
			if dp[i] > -dp[i+1]+1 {
				dp[i] = dp[i+1] - 1
			} else if dp[i] == -dp[i+1]+1 {
				dp[i] = 0
			}
		}

		if dp[i] > 0 {
			res[i] = 'R'
		} else if dp[i] < 0 {
			res[i] = 'L'
		} else {
			res[i] = '.'
		}
	}
	return string(res)
}

func PushDominoes(dominoes string) string {
	return pushDominoes(dominoes)
}
