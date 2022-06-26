func makeMapFromString(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	return m
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	m2 := makeMapFromString(p)
	res := make([]int, 0)

	for i := 0; i <= len(s)-len(p); i++ {
		m1 := makeMapFromString(s[i : i+len(p)])
		if reflect.DeepEqual(m1, m2) {
			res = append(res, i)
		}
	}
	return res
}
