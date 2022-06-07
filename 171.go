func titleToNumber(columnTitle string) int {
	res := 0
	n := len(columnTitle)
	k := 1
	for i := n - 1; i >= 0; i-- {
		res = res + (int(columnTitle[i]-'A') + 1) * k
		k *= 26
	}
	return res
}
