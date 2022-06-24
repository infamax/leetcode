func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := make(map[rune]int)

	for _, v := range s {
		m[v]++
	}

	res := 0
    countEven := 0
	for _, value := range m {
		if value%2 == 0 {
			res += value
		} else {
            countEven++
			res += value
			res--
		}
	}
    
    if countEven > 0 {
	    res++
    }
	return res
}
