
func makeMapSymbols(word string) map[rune]int {
	m := make(map[rune]int)
	for _, s := range word {
		m[s]++
	}
	return m
}

func canConstruct(ransomNote string, magazine string) bool {
	m1 := makeMapSymbols(ransomNote)
	m2 := makeMapSymbols(magazine)
	for key, value := range m1 {
		v, ok := m2[key]
		if !ok {
			return false
		}
		if v < value {
			return false
		}
	}
	return true
}
