func countVowelStrings(n int) int {
	dp := []int{1, 1, 1, 1, 1}
	for i := 2; i <= n; i++ {
		for j := 1; j < 5; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[0] + dp[1] + dp[2] + dp[3] + dp[4]
}
