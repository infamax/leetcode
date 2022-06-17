
func makeMapByString(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

// s2 contains permutation s1
func check(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	m1 := makeMapByString(s1)
	l := 0
	r := len(s1) - 1
	for r < len(s2) {
		m2 := makeMapByString(s2[l : r+1])
		if reflect.DeepEqual(m1, m2) {
			return true
		}
		l++
        r++
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	return check(s1, s2)
}
