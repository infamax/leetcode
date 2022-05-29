func makeMapWords(s string) map[rune]bool {
	m := make(map[rune]bool)
	for _, ch := range s {
		m[ch] = true
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(words []string) int {
	res := 0
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			m1 := makeMapWords(words[i])
			m2 := makeMapWords(words[j])
			fl := true
			for ch := 'a'; ch <= 'z'; ch++ {
				_, ok1 := m1[ch]
				_, ok2 := m2[ch]
				if ok1 && ok2 {
					fl = false
				}
			}
			if fl {
				res = max(res, len(words[i]) * len(words[j]))
			}
		}
	}
	return res
}
