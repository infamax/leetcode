func fib(n int) int {
	if n == 0 {
		return 0
	}
	
	if n == 1 {
		return 1
	}
	
	a := 0
	b := 1
	res := a + b
	for i := 2; i < n; i++ {
		a = b
		b = res
		res = a + b
	}
	return res
}
