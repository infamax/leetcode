func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]rune)

	for i, v := range s {
		symbol, ok := m[v]
		if ok {
			if symbol != rune(t[i]) {
				return false
			}
		}
		m[v] = rune(t[i])
		
	}
	return true
}
