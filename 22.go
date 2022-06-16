func generateParenthesis(n int) []string {
	var gen func(n int, co int, cc int, res string)
	var ans []string
	gen = func(n int, co int, cc int, res string) {
		if co+cc == n {
			ans = append(ans, res)
			return
		}
		
		if co < n / 2 {
			gen(n, co + 1, cc, res + "(")
		}
		
		if cc < co {
			gen(n, co, cc + 1, res + ")")
		}
	}
	gen(2*n, 0, 0, "")
	return ans
}
