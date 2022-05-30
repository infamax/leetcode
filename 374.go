func guessNumber(n int) int {
	l := 0
	r := n + 1
	for r-l > 1 {
		m := (l + r) / 2
		if guess(m) == 0 {
			return m
		}

		if guess(m) == -1 {
			r = m
		} else {
			l = m
		}
	}
	if guess(l) == 0 {
		return l
	}
	return r
}
