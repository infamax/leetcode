func sumSquares(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}

func isHappy(n int) bool {
	for n > 4 {
		n = sumSquares(n)
	}
	if n >= 2 && n <= 4 {
		return false
	}
	return true
}
