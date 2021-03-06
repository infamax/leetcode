func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	
	if n == 1 || n == 2 {
		return 1
	}
	
	a := 0
	b := 1
	c := 1
	res := 0
	for i := 3; i <= n; i++ {
		res = a + b + c
		a = b
		b = c
		c = res
	}
	return res
}
