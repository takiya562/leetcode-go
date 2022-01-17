package leetcode

func countVowelPermutation(n int) int {
	dp := make([]int, 5)
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		pre := dp
		dp = make([]int, 5)
		dp[0] = (pre[1] + pre[2] + pre[4]) % 1000000007
		dp[1] = (pre[0] + pre[2]) % 1000000007
		dp[2] = (pre[1] + pre[3]) % 1000000007
		dp[3] = pre[2] % 1000000007
		dp[4] = (pre[2] + pre[3]) % 1000000007
	}
	return (dp[0] + dp[1] + dp[2] + dp[3] + dp[4]) % 1000000007
}
