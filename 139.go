func contains(word string, wordDict []string) bool {
	for _, item := range wordDict {
		if item == word {
			return true
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s) + 1, len(s) + 1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := i - 1; j >= 0; j-- {
            if dp[j] && contains(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
