func sumDigits(num int) int {
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}

func addDigits(num int) int {
	if num == 0 {
		return num
	}
	
	for num > 9 {
		num = sumDigits(num)
	}
	return num
}
