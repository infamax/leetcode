
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	l := 0
	r := num
	for r-l > 1 {
		m := (l + r) / 2
		if m*m < num {
			l = m
		} else {
			r = m
		}
	}
	if r * r == num {
		return true
	}
	return false
}
