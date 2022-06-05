func fizzBuzz(n int) []string {
	res := make([]string, n, n)
	for i := 0; i < n; i++ {
		number := i + 1
		if number%3 != 0 && number%5 != 0 {
			res[i] = strconv.Itoa(number)
		} else if number%3 == 0 && number%5 != 0 {
			res[i] = "Fizz"
		} else if number%3 != 0 && number%5 == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = "FizzBuzz"
		}
	}
	return res
}
