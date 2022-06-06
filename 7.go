const inf = math.MaxInt32

func countRanks(x int) int {
	k := 0
	for x > 0 {
		k++
		x /= 10
	}
	return k
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	fl := false
	if x < 0 {
		fl = true
		x = -x
	}

	res := 0
	k := countRanks(x) - 1
	for x > 0 {
		num := x % 10
		x /= 10

		res += num * int(math.Pow10(k))
		k--
	}

	if res > inf {
		return 0
	}
	
	if fl {
		res = -res
	}
	
	return res
}
