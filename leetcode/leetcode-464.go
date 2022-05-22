package leetcode

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if desiredTotal == 0 {
		return true
	}
	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}
	dp := make([]int, 1<<maxChoosableInteger)
	var dfs func(int, int) int
	dfs = func(state, total int) int {
		if dp[state] != 0 {
			return dp[state]
		}
		for i := 0; i < maxChoosableInteger; i++ {
			if (state & (1 << i)) != 0 {
				continue
			}
			if total+i+1 >= desiredTotal {
				dp[state] = 1
				return 1
			}
			if dfs(state|(1<<i), total+i+1) == -1 {
				dp[state] = 1
				return 1
			}
		}
		dp[state] = -1
		return -1
	}
	return dfs(0, 0) == 1
}

func CanIWin(maxChoosableInteger int, desiredTotal int) bool {
	return canIWin(maxChoosableInteger, desiredTotal)
}
